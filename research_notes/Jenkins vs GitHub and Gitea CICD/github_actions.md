# GitHub Actions as a CI/CD Platform: State as of November 2026

Research date: 2026-11-02. Figures are from the dated source cited. Vendor and blog claims are labelled as such.

## What are the current runner options and workflow features?

### Takeaway
Actions offers standard hosted, larger hosted, self-hosted, and Kubernetes-autoscaled (ARC with runner scale sets) runners. Its composition model is mature but tightly bounded (256-job matrix, 10 levels and 50 reusable workflows per run). The December 2025 re-architecture and ARC 0.14/0.15 are the main recent platform changes.

### Cited Findings

**Runner types**
- Standard hosted runners: Linux 1-core x64, Linux 2-core x64 and arm64, Windows 2-core x64 and arm64, macOS 3- or 4-core (M1 or Intel) — [GitHub Docs: runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)
- Larger runners: x64 Linux up to 96-core, x64 Windows up to 96-core, arm64 Linux and Windows up to 64-core, macOS 12-core x64 and 5-core M2 Pro arm64, and 4-core GPU on Linux and Windows. They are available only on Team or Enterprise Cloud and are billed even for public repos. Static IPs carry no extra cost. Custom images work only on larger runners, with image storage billed separately — [GitHub Docs: runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)
- Larger runners support static IPs (10 per enterprise or org, increasable) and vnet injection. Subnets need a 30% buffer above peak concurrency, and Azure reserves 5 IPs per subnet — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)
- A runner scale set is a group of homogeneous runners that can be assigned jobs. Workflows target it by name or label. A set belongs to one runner group at a time, and autoscalers such as ARC control how many active runners it owns — [GitHub Docs: runner scale sets](https://docs.github.com/en/actions/concepts/runners/runner-scale-sets)
- ARC 0.14.0 (March 19, 2026) added multi-label scale sets, moved the controller to the `actions/scaleset` client library, and introduced experimental Helm chart rewrites — [GitHub Changelog](https://github.blog/changelog/2026-03-19-actions-runner-controller-release-0-14-0/)
- ARC 0.15.0 (October 1, 2026) focused on reliability and scale: in-place resource updates on patch upgrades, configurable shutdown grace periods, `max-concurrent-reconciles` flags, and configurable listener QPS and burst — [GitHub Changelog](https://github.blog/changelog/2026-10-01-actions-runner-controller-release-0-15-0/). A 0.15.1 patch followed — [ARC releases](https://github.com/actions/actions-runner-controller/releases)

**Limits**
- Run and job durations: workflow run 35 days including waiting; environment approval wait 30 days; hosted job 6 hours; self-hosted job 5 days; self-hosted queue 24 hours — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)
- Matrix: 256 jobs per run. Re-runs: 50 per run. Workflow file: 500 KB. Trigger events: 1,500 per 10 seconds per repo. Queued runs: 500 per 10 seconds. `queue: max` concurrency groups: 100 jobs or runs per group — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)
- Self-hosted: 10,000 runners per runner group; 1,500 registrations per 5 minutes — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)
- Concurrent jobs on standard runners: Free 20, Pro 40, Team 60, Enterprise 500. macOS: 5 (Enterprise 50). Larger runners: 1,000 total and 100 GPU. Concurrency can be raised by support ticket — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)
- `GITHUB_TOKEN` API limit: 1,000 per hour per repo (15,000 for Enterprise Cloud resources). Docker Hub pulls from self-hosted runners are always rate limited, while public-image pulls from hosted runners are not — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits)

**Workflow composition and identity**
- Reusable workflows: up to 10 levels of nesting and 50 unique reusable workflows from one workflow file, counting the entire nested tree. They are called from a job rather than a step — [GitHub Docs: reusing workflows](https://docs.github.com/en/actions/reference/workflows-and-actions/reusing-workflow-configurations)
- These limits were raised as of December 2025, alongside `workflow_dispatch` inputs going from 10 to 25 — [GitHub Blog, Dec 11, 2025](https://github.blog/news-insights/product-news/lets-talk-about-github-actions/)
- OIDC: a JWT is issued per job and validated against trust conditions on the cloud role, removing the need for long-lived cloud secrets. Named providers are AWS, Azure, GCP, and HashiCorp Vault. Claims include `sub`, `ref`, `sha`, `repository`, `environment`, and `job_workflow_ref` — [GitHub Docs: OIDC](https://docs.github.com/en/actions/concepts/security/openid-connect)
- Repositories created after July 15, 2026 use an immutable OIDC subject format including owner and repository IDs, which is not available on GitHub Enterprise Server — [GitHub Docs: OIDC](https://docs.github.com/en/actions/concepts/security/openid-connect)

**Architecture and scale**
- GitHub re-architected the core backend for jobs and runners, with all jobs on the new architecture since August 2025. Volume went from about 23M jobs/day (early 2024) to 71M jobs/day (Dec 2025). Enterprises can start 7x more jobs per minute. The work delayed feature development — [GitHub Blog, Dec 11, 2025](https://github.blog/news-insights/product-news/lets-talk-about-github-actions/)
- Public and open source projects used 11.5B Actions minutes in 2025, up 35% year over year — [GitHub Blog, Dec 11, 2025](https://github.blog/news-insights/product-news/lets-talk-about-github-actions/)
- The same post promised for Q1 2026 scheduled-job time zones, run IDs returned from dispatch, and an expression conditional, plus parallel steps targeted before mid-2026 — [GitHub Blog, Dec 11, 2025](https://github.blog/news-insights/product-news/lets-talk-about-github-actions/)

### Inferences
- 71M jobs/day against a stated 10x design goal, combined with later "Actions growth" capacity admissions (see reliability), suggests the re-platform relieved but did not end scaling pressure.
- Self-hosted users still depend on GitHub's control plane. The July 19–20, 2026 certificate incident delayed self-hosted and scale-set jobs, so self-hosting does not isolate a team from GitHub outages.

### Gaps
- Did not verify whether parallel steps and the other promised Q1 2026 items shipped. The October 2026 changelog review found no such entries, but earlier months were not checked.
- No primary source collected on composite action nesting limits or a composite versus reusable comparison, as the fetched docs page did not cover them.
- Did not confirm the scale set numeric limits or the state of a standalone scale set client outside ARC beyond the `actions/scaleset` library mention.

## What is current pricing, and what happened with the 2025–2026 changes?

### Takeaway
The hosted-runner cut (up to 39%) took effect January 1, 2026. The $0.002/min self-hosted charge announced for March 1, 2026 never took effect: GitHub postponed it, not cancelled it, and current billing docs say self-hosted usage is free. No new date was found as of November 2, 2026.

### Cited Findings

**Timeline and status**
- December 16, 2025: GitHub announced hosted-runner price cuts of up to 39% from January 1, 2026, and a "$0.002 per minute GitHub Actions cloud platform charge" for self-hosted usage from March 1, 2026, counting toward plan minutes. It said runner usage in public repos stays free — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- Reversal text: "We're postponing the announced billing change for self-hosted GitHub Actions to take time to re-evaluate our approach." No new date is given, and the update itself on the fetched page is undated. Hosted cuts proceeded and free quotas stayed the same — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- Current status: billing docs state that self-hosted usage "remains free independently and does not reduce the included compute usage" and mention no upcoming charges — [GitHub Docs: Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions)
- GitHub's claimed impact: "96% of customers will see no change to their bill"; about 0.09% of individual private-repo Actions users would see an increase, with a median under $2/month — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- Enterprise Server: "There will be no changes in price structure for GitHub Enterprise Server customers" — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- Third parties independently describe the fee as postponed or suspended rather than cancelled, one as of October 2026 — [RunsOn blog](https://runs-on.com/blog/github-self-hosted-runner-fee-2026/); [Plainwork Digital](https://plainworkdigital.com/d20261012a9/github-actions-self-hosted-runner-cost)

**Current rates (USD per minute)**
- Standard: Linux 1-core $0.002, Linux 2-core x64 $0.006, Linux 2-core arm64 $0.005, Windows 2-core $0.010, macOS 3/4-core $0.062 — [GitHub Docs: runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)
- Larger x64 Linux: 4-core $0.012, 8-core $0.022, 16-core $0.042, 32-core $0.082, 64-core $0.162, 96-core $0.252. Larger arm64 Linux: 4-core $0.008 to 64-core $0.098. GPU: Linux $0.052, Windows $0.102. macOS: 12-core $0.077, M2 Pro $0.102 — [GitHub Docs: runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)
- Each job is rounded up to the nearest whole minute — [GitHub Docs: runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)

**Included usage and storage**
- Included compute is now stated in dollars per month: Free and Free for orgs $12; Pro and Team $18; Enterprise Cloud $300 (50,000 minutes at $0.006) — [GitHub Docs: Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions)
- Standard hosted runners are free in public repos. Artifacts and Packages storage is $0.25/GB-month. Cache and custom image storage are $0.07/GB-month — [GitHub Docs: Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions)
- Cache: 10 GB per repo is free, and repos can exceed it pay-as-you-go since November 2025 — [GitHub Changelog, Nov 20, 2025](https://github.blog/changelog/2025-11-20-github-actions-cache-size-can-now-exceed-10-gb-per-repository/). If a budget is exceeded, the cache becomes read-only — [GitHub Docs: dependency caching](https://docs.github.com/en/actions/reference/workflows-and-actions/dependency-caching)

### Inferences
- The dollar allowances match the old minute quotas at $0.006/min: $12 = 2,000 minutes, $18 = 3,000, $300 = 50,000. This is consistent with GitHub's statement that quotas stayed the same.
- A self-hosted platform fee remains a live planning risk. The reversal text is "postpone" and "re-evaluate," not "cancel." At $0.002/min, 1M private-repo self-hosted minutes per month would be about $2,000, before any included-minute offset (arithmetic based on the announced rate).
- BYO-compute vendors running on the self-hosted protocol would have been exposed to the same fee, so it should be stress-tested in any cost comparison involving them.

### Gaps
- The exact postponement date could not be sourced from a primary page. Third parties say it came within days of December 16, 2025; the Socket write-up returned HTTP 403 and could not be read.
- Pre-2026 per-minute rates were not collected from a primary source, so per-SKU reduction percentages are not listed here.
- GHES-specific Actions licensing beyond the "no changes" line was not found.

## What is the security record in 2025–2026, and how has the platform responded?

### Takeaway
Three high-impact episodes (tj-actions/reviewdog in March 2025, Shai-Hulud waves in September and November 2025, and Trivy in March 2026) exposed mutable tags, secret access at runtime, and fork-triggered workflows. GitHub's response accelerated in 2026 and has now shipped concrete controls, with a key enforcement date of December 2, 2026 for `pull_request_target` on public repos.

### Cited Findings

**Incidents**
- tj-actions/changed-files, CVE-2025-30066: all versions used from 2025-03-12 00:00 UTC to 2025-03-15 12:00 UTC were affected; patched in v46.0.1. Reviewdog/action-setup@v1, CVE-2025-30154: March 11, 2025, 18:42–20:31 UTC. CISA's alert is dated March 18, 2025, revised March 26, and both CVEs were added to KEV — [CISA](https://www.cisa.gov/news-events/alerts/2025/03/18/supply-chain-compromise-third-party-github-action-cve-2025-30066)
- Root cause and scale: Wiz linked the reviewdog compromise to the theft of the tj-actions-bot PAT. Palo Alto found only 218 repositories actually leaked secrets, mostly short-lived tokens, and that the attack initially targeted Coinbase — [SecurityWeek](https://www.securityweek.com/impact-root-cause-of-github-actions-supply-chain-hack-revealed/). One source reported over 23,000 dependent repositories — [InfoWorld](https://www.infoworld.com/article/3849245/github-suffers-a-cascading-supply-chain-attack-compromising-ci-cd-secrets.html)
- Shai-Hulud, first wave (September 14, 2025): worm via compromised maintainer accounts; pushed `shai-hulud-workflow.yml` Actions workflows to exfiltrate secrets; GitHub removed 500+ packages — [GitHub Blog](https://github.blog/security/supply-chain-security/our-plan-for-a-more-secure-npm-supply-chain/); [ADTmag](https://adtmag.com/blogs/watersworks/2025/10/the-worm-that-ate-javascript.aspx)
- Shai-Hulud 2.0 (November 24–27, 2025): about 700–800 npm packages and over 25,000 repos, with internally inconsistent counts in the source. It registered infected machines as self-hosted runners named `SHA1HULUD` and added a `discussion.yaml` injection backdoor. Wiz found no in-the-wild use of the backdoor but confirmed it works. A `formatter` workflow dumped `toJSON(secrets)` as an artifact, then deleted the workflow and branch — [Wiz](https://www.wiz.io/blog/shai-hulud-2-0-ongoing-supply-chain-attack)
- Trivy, CVE-2026-33634 (March 19–20, 2026): 76 of 77 trivy-action tags and all 7 setup-trivy tags force-pushed. Exposure: trivy-action about 12 hours, setup-trivy about 4 hours. Root cause was non-atomic credential rotation after a late-February intrusion — [Aqua GHSA](https://github.com/aquasecurity/trivy/security/advisories/GHSA-69fq-xp46-6x23). CISA added it to KEV on March 26, 2026 — [NVD](https://nvd.nist.gov/vuln/detail/CVE-2026-33634); attribution to TeamPCP comes from vendors — [Legit Security](https://www.legitsecurity.com/blog/the-trivy-supply-chain-compromise-what-happened-and-playbooks-to-respond)
- Trivy detail for pinning policy: v0.35.0 and v0.69.3 were protected by immutable releases enabled March 3–4. However, even a safe SHA of trivy-action from before April 9, 2025 could still pull malicious setup-trivy during the window, because a nested dependency was unpinned — [Aqua GHSA](https://github.com/aquasecurity/trivy/security/advisories/GHSA-69fq-xp46-6x23)

**Platform response (chronological)**
- August 15, 2025: Actions policy gained blocklisting and required full-length SHA pinning; unpinned action use then fails the workflow — [GitHub Changelog](https://github.blog/changelog/2025-08-15-github-actions-policy-now-supports-blocking-and-sha-pinning-actions/)
- December 8, 2025: `pull_request_target` runs take their workflow file from the default branch, and environment rules evaluate against the default branch — [GitHub Changelog](https://github.blog/changelog/2025-11-07-actions-pull_request_target-and-environment-branch-protections-changes/)
- March 26, 2026: 2026 security roadmap published, citing tj-actions, Nx, and trivy-action. Items: dependency locking, policy-driven execution, scoped secrets, secrets permission, Actions Data Stream, native egress firewall, in relative 3–9 month windows — [GitHub Blog](https://github.blog/news-insights/product-news/whats-coming-to-our-github-actions-2026-security-roadmap/)
- June 18, 2026: `actions/checkout` v7 blocks fork PR code in `pull_request_target` and `workflow_run` (for `pull_request*` events) by default. Backport to supported majors was enforced July 20, 2026. Pinned SHA, minor, and patch versions do not get the change. Not covered: fork code fetched by `git` or `gh` in `run` blocks, `issue_comment`, and unrelated third-party repos. Opt-out: `allow-unsafe-pr-checkout` — [GitHub Changelog](https://github.blog/changelog/2026-06-18-safer-pull_request_target-defaults-for-github-actions-checkout/)
- June 2026: read-only Actions cache for untrusted triggers — [GitHub Changelog](https://github.blog/changelog/2026-06-26-read-only-actions-cache-for-untrusted-triggers/). Also workflow trigger controls, npm v12 disabling install scripts by default, and credential revocation tools. July 2026: Dependabot default 3-day cooldown for version updates — [GitHub Blog, Jul 28, 2026](https://github.blog/security/supply-chain-security/disrupting-supply-chain-attacks-on-npm-and-github-actions/)
- July 28, 2026: Actions network firewall is in technical preview, logging outbound traffic; egress restrictions and policies are future work — [GitHub Blog](https://github.blog/security/supply-chain-security/disrupting-supply-chain-attacks-on-npm-and-github-actions/)
- September 10, 2026: `cache-mode` control for Actions cache access — [GitHub Changelog](https://github.blog/changelog/2026-09-10-control-github-actions-cache-access-with-cache-mode/)
- October 14 and 28, 2026: workflow lockfiles reached public preview and then GA on Enterprise Cloud, data residency, Team, and Free. `.github/workflows/actions.lock` pins actions and their transitive dependencies, generated with the `gh actions-lock` extension. Tags in YAML stay readable while locked commits run. Dependabot regenerates entries. A "Require lockfile" policy enforces adoption — [GitHub Changelog: GA](https://github.blog/changelog/2026-10-28-github-actions-workflow-lockfiles-are-generally-available/); [preview](https://github.blog/changelog/2026-10-14-github-actions-workflow-lockfiles-in-public-preview/)
- December 2, 2026 (pending): a default policy will block `pull_request_target` in public repos without another applicable event policy. It is currently in evaluate mode. Private and internal repos and repos with an applicable policy are unaffected. Opt-out is an explicit allow policy — [GitHub Docs](https://docs.github.com/en/actions/reference/security/securely-using-pull_request_target); same date on [Enterprise Cloud docs](https://docs.github.com/en/enterprise-cloud@latest/actions/how-tos/administer/control-workflow-execution)

**Pinning and immutability**
- Practitioner view: immutable releases have been GA since October 2025, but no Actions policy treats an immutable release as immutable; only full-length SHA pinning can be enforced — [Jesse Houwing, DEV](https://dev.to/jessehouwing/keep-the-creepy-crawlies-out-of-you-workflows-securing-github-actions-in-2026-2np0). This may be partly dated by the October 2026 lockfile GA.
- GitHub's roadmap acknowledges commit SHAs are hard to manage at scale and plans hardened publishing and stricter release requirements, with no timeline given — [GitHub Blog](https://github.blog/news-insights/product-news/whats-coming-to-our-github-actions-2026-security-roadmap/)
- A community feedback thread notes Dependabot cannot update npm dependencies of actions and that immutable releases do not solve this; a GitHub comment says it plans to address that — [GitHub Discussion #190621](https://github.com/orgs/community/discussions/190621)

### Inferences
- The incidents map to distinct fixes: tag repointing to lockfiles and SHA policy; transitive action dependencies to lockfiles (Trivy shows SHA pinning alone is insufficient); fork code execution to checkout v7 and the default policy; egress exfiltration to the firewall, which remains logging-only as of its last documented status.
- Most protections are opt-in except the checkout backport and the upcoming December 2 default policy. Workflows pinned by SHA do not get the checkout protection until upgraded, a small but real pinning trade-off.
- Teams with public repos and `pull_request_target` workflows should check policy insights before December 2, 2026.

### Gaps
- Status of the original "immutable actions" product (GHCR-backed immutable action packages) was not found.
- Scoped secrets, secrets permission, Actions Data Stream, and execution rulesets: no launch confirmation located beyond June's trigger controls. The October changelog review found no entries.
- The lockfile GA post does not mention GitHub Enterprise Server, so GHES availability is unconfirmed.
- GitHub's post on supply chain attacks names no specific campaign, so the direct link between Shai-Hulud and these controls is inference. No StepSecurity primary was read; findings rely on CISA, Wiz, the GHSA, and SecurityWeek's account of Palo Alto's findings.

## What is the reliability record, and how has the Azure migration affected it?

### Takeaway
Reliability is the platform's weakest point in 2026. Actions had the most outages of any GitHub service in one May 2025–April 2026 tally, and multi-hour Actions incidents recurred through October 2026. GitHub attributes much of this to AI-driven load and capacity limits, while the Azure migration is both the capacity fix and a documented source of risk.

### Cited Findings

**Volume**
- Third-party tally, May 2025–April 2026: 257 GitHub incidents (48 major); Actions 57 outages (16 major), the most of any service. Actions peaked at 8 in January and February 2026. Root causes: Capacity 83, Deployment 71, External Dependencies 31. GitHub's count far exceeds GitLab's 132 and Bitbucket's 27, not normalized for traffic — [IncidentHub, Apr 30, 2026](https://blog.incidenthub.cloud/github-reliability-outage-history-2025-2026)
- Counts differ by method: availability reports listed 10 incidents in April and 9 in May 2026, versus 26 and 23 on the incident history page. An unofficial tracker reported monthly availability of 78.33% (April), 93.86% (May), and 88.39% (June to date), against about 99.9% on the official status page — [The Register, Jun 12, 2026](https://www.theregister.com/software/2026/06/12/github-outages-persist-as-ai-coding-drives-traffic-surge/5255125)

**Actions incidents in 2026 (GitHub's own reports)**
- February 2: security policies wrongly applied to internal storage accounts in all regions blocked VM metadata access and halted hosted runner lifecycle. March 5: Redis failover left the Actions job orchestration cluster with no writable primary — [GitHub Blog, Mar 11, 2026](https://github.blog/news-insights/company-news/addressing-githubs-recent-availability-issues-2/)
- July 8 (7h 4m): multi-service outage; Actions runs queued or rejected. July 9 (9h 18m): about 8% of runs delayed over 5 minutes and about 2% failed to start. July 19–20 (5h 10m): expired SSL certificate; 78.99% of larger-runner jobs, 29.8% of scale-set jobs, and 8.7% of self-hosted jobs waited over 5 minutes. July 25: up to 60% of runs failed with infrastructure errors at peak — [GitHub, July report](https://github.blog/news-insights/company-news/github-availability-report-july-2026/)
- August 6 (10h 42m by header, about 9h by summary): hosted and self-hosted runners; at least 74 orgs. August 26 (2h 50m): more than one in five run starts per minute failed or was badly delayed. GitHub: "shared infrastructure services have not kept up with our month-over-month actions growth and peak load" — [GitHub, August report](https://github.blog/news-insights/company-news/github-availability-report-august-2026/)
- September 13: about 4% of Actions workflows had intermittent failures in a shared-database incident — [GitHub, September report](https://github.blog/news-insights/company-news/github-availability-report-september-2026/)
- October 2026: the status feed shows 12 entries naming Actions or Actions-related problems, including October 1 job delays (about 3h41m), October 5 (about 4h01m), and October 19 (1h28m) — [GitHub Status feed](https://www.githubstatus.com/history.atom)

**Azure migration and capacity**
- Planned capacity increase went from 10x (October 2025) to 30x (February 2026). Commits rose from 1B in all of last year (reported) to 1.4B per month — [The Register](https://www.theregister.com/software/2026/06/12/github-outages-persist-as-ai-coding-drives-traffic-surge/5255125)
- Migration progress; these are different metrics, not a single series:
  - March 2026: 12.5% of all traffic from Azure Central US, with a 50% target by July — [GitHub Blog](https://github.blog/news-insights/company-news/addressing-githubs-recent-availability-issues-2/)
  - June 2026 statement: 40% of monolith traffic (8% in February), Git 30% — [The Register](https://www.theregister.com/software/2026/06/12/github-outages-persist-as-ai-coding-drives-traffic-surge/5255125)
  - July 28: monolith reads peaked at 52.75%; Git 47% — [July report](https://github.blog/news-insights/company-news/github-availability-report-july-2026/)
  - October 7: core app reads 65.4%, Git reads 48.3%, stream processing 80.2%; 120 of 160 production services deployed and 13 of 51 primary MySQL databases moved — [September report](https://github.blog/news-insights/company-news/github-availability-report-september-2026/)
- First production MySQL primary run from Azure on August 11, 2026 — [August report](https://github.blog/news-insights/company-news/github-availability-report-august-2026/)
- Actions-specific: GitHub is accelerating the move of Actions to Azure, though the launch service bridging the monolith to Actions remained in GitHub datacenters as of July — [July report](https://github.blog/news-insights/company-news/github-availability-report-july-2026/)
- Actions containment: job routing moved 33% of jobs to spare capacity, cutting peak cache CPU from 98% to 80%, about three months of headroom and "not a permanent fix" — [August report](https://github.blog/news-insights/company-news/github-availability-report-august-2026/). Later changes cut Redis operations about 28% and CPU about 30% — [September report](https://github.blog/news-insights/company-news/github-availability-report-september-2026/)
- Migration-caused issues: September 11 pause of Central US ramps over uneven availability-zone capacity (resumed September 18); a planned failover disrupted customers in that region; the September 30 goal was missed — [September report](https://github.blog/news-insights/company-news/github-availability-report-september-2026/). Migration load saturating database connections appears in the [May report](https://github.blog/news-insights/company-news/github-availability-report-may-2026/) per search summary only.
- Stated priority: "availability, then capacity, then features" — [August report](https://github.blog/news-insights/company-news/github-availability-report-august-2026/)
- Some Actions incidents (October 29, 2025 and February 2, 2026) were linked to Azure as a provider — [IncidentHub](https://blog.incidenthub.cloud/github-reliability-outage-history-2025-2026)

### Inferences
- Causes are mostly control-plane (databases, Redis, service mesh, job assignment), so faster third-party or self-hosted compute does not protect against them.
- The March target of 50% of all traffic by July appears only partly met: monolith reads passed 50% in late July, but write traffic was just beginning and the Git share was lower. The metric definitions differ, so this cannot be stated definitively.
- GitHub's feature velocity is explicitly subordinate to availability, which may explain why roadmap items slip.

### Gaps
- No official numeric SLA attainment for Actions in 2026 was collected. Enterprise SLA terms were not researched.
- AWS emergency capacity was reported by one low-confidence site ([tech-insider](https://tech-insider.org/microsoft-aws-github-outages-2026/)); The Register's article does not mention it. Treat it as unverified.
- The October availability report had not been found as of research date. October figures are feed entries, not root-cause analysis.

## What limits and pain points do senior practitioners report?

### Takeaway
Consistent complaints are YAML and expression fragility, weak local and interactive debugging, a poor log viewer, opaque caching, coarse concurrency, and supply chain risk. A sizeable vendor market exists to replace the compute layer, but it does not remove YAML or control-plane dependence.

### Cited Findings
- Ian Duncan (Feb 5, 2026): log viewer needs many clicks and can crash browsers on large logs; expression language has odd interpolation and conditionals, with errors surfacing only after slow runner spin-up; marketplace actions are opaque code with access to secrets; `actions/cache` has confusing keys, silent misses, and opaque eviction; concurrency is coarse; secrets cannot be used in `if`. He considers it fine for small teams and recommends Buildkite for larger — [Ian Duncan](https://www.iankduncan.com/engineering/2026-02-05-github-actions-killing-your-team/) (opinion)
- Duncan reads the third-party runner market (Blacksmith, Namespace, Actuated, RunsOn, BuildJet) as proof that default runners are inadequate — [Ian Duncan](https://www.iankduncan.com/engineering/2026-02-05-github-actions-killing-your-team/) (opinion)
- Interactive debugging is not natively available, a request open since 2019 by this author's account, and standard YAML linters miss semantically invalid keys — [htek.dev](https://htek.dev/articles/github-actions-debugging-guide)
- Community discussion echoes these complaints and suggests moving logic out of YAML or self-hosting — [Hacker News, 2025](https://news.ycombinator.com/item?id=43419701)
- Local execution: `act` does not reproduce `GITHUB_TOKEN` permissions, and limitations are most visible around secrets, cache, service containers, and permissions — [OneUptime, Feb 2026](https://oneuptime.com/blog/post/2026-02-08-how-to-run-act-to-test-github-actions-locally-in-docker/view). This is secondary; the `act` repo was not read directly.
- Cache: 10 GB free, default 7-day retention, LRU eviction — [GitHub Docs: dependency caching](https://docs.github.com/en/actions/reference/workflows-and-actions/dependency-caching). 200 uploads and 1,500 downloads per minute per repo — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits); upload rate limit dated January 2026 — [GitHub Changelog](https://github.blog/changelog/2026-01-16-rate-limiting-for-actions-cache-entries/)
- High-profile departure: Zig moved its main repo to Codeberg in late November 2025, citing Actions bugs and unpredictable scheduling, alongside its AI policy concerns — [heise](https://www.heise.de/en/news/Zig-turns-its-back-on-GitHub-Frustration-over-Actions-and-Microsoft-s-AI-course-11100039.html?view=print); [DevClass](https://devclass.com/2025/11/27/zig-project-ditches-github-for-codeberg-but-move-could-be-costly/)

**Third-party runners (largely vendor data; verify before use)**
- RunsOn: flat annual license, with ephemeral EC2 runners in your own AWS account and compute billed by AWS at spot prices; lists roughly $0.0006–$0.0045 per minute for small to mid-size runners — [RunsOn pricing](https://runs-on.com/pricing/)
- Depot: Developer plan 2,000 minutes per month, Startup 20,000; extra usage $0.006/min for a 2-vCPU Linux runner — [Depot calculator](https://depot.dev/github-actions-price-calculator); [Depot docs](https://depot.dev/docs/github-actions/overview)
- Blacksmith: its own posts conflict, with $0.004/min for 2 vCPU in one and $0.008/min in an older one — [Blacksmith](https://info.blacksmith.sh/task/blog/github-actions-runner-services-cost-effective-scaling); [Blacksmith](https://www.blacksmith.sh/blog/how-to-reduce-spend-in-github-actions)

**Why teams switch tools generally (not Actions-specific)**
- 62% of developers have no plans to switch CI/CD tool; those who do cite scaling costs (24%), maintenance effort (23%), consolidation (22%), and complex configuration (21%) — [JetBrains State of CI/CD 2026](https://www.jetbrains.com/pages/state-of-cicd2026/)

### Inferences
- The vendor layer solves performance and unit cost, but all of these still run through GitHub's workflow syntax and job assignment. Lock-in lives mainly in workflow YAML, marketplace actions, and OIDC or environment configuration.
- Moving logic into scripts or Make/Dagger-like tools is the common practitioner mitigation for debugging and portability.
- The strongest published critiques are from individuals advocating alternatives, so they indicate real issues but not frequency.

### Gaps
- No independent, current benchmark of Blacksmith, Depot, and RunsOn was found; available comparisons are vendor-run or vendor-affiliated. Current official Blacksmith pricing page and the RunsOn license price were not verified.
- No primary quantitative survey on the prevalence of specific Actions pain points was located.
- Docs conflict on cache: the limits page lists 10 GB as non-increasable storage, whereas the changelog and billing docs describe paid expansion. The paid path is treated as current.

## What are its real strengths?

### Takeaway
Actions' strengths are distribution and integration: it leads survey use, is built into the repo workflow, needs little operation for hosted runners, offers OIDC and policy controls in one platform, and is becoming the execution layer for GitHub's agent workflows.

### Cited Findings
- Adoption: 47% organizational use in 2026 (41% in 2025), ahead of GitLab CI at 32% (34%), Jenkins at 28% (31%), and Azure DevOps Server at 15% (14%). n=598; multi-select, not market share — [JetBrains State of CI/CD 2026](https://www.jetbrains.com/pages/state-of-cicd2026/). Other blogs quote 33–40% or 60%+ with weak sourcing; prefer the JetBrains primary.
- Counter-trend: on-premises installations reached 31%, a five-year high, and fully managed cloud CI is down seven points from its 2022 peak — [JetBrains State of CI/CD 2026](https://www.jetbrains.com/pages/state-of-cicd2026/)
- Open source economics: standard hosted runners are free in public repos — [GitHub Docs: Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions). 11.5B public and open source minutes in 2025 — [GitHub Blog](https://github.blog/news-insights/product-news/lets-talk-about-github-actions/)
- Low operating burden and elasticity: hosted concurrency of 500 standard or 1,000 larger jobs on Enterprise, with hardware from 1 to 96 cores and GPU, without managing a controller — [GitHub Docs: limits](https://docs.github.com/en/actions/reference/limits); [runner pricing](https://docs.github.com/en/billing/reference/actions-runner-pricing)
- Integrated identity: per-job OIDC to AWS, Azure, GCP, and Vault — [GitHub Docs: OIDC](https://docs.github.com/en/actions/concepts/security/openid-connect)
- Integrated governance: SHA pinning, lockfile, and trigger policies at enterprise, org, or repo level; Dependabot understands lockfiles — [GitHub Changelog](https://github.blog/changelog/2026-10-28-github-actions-workflow-lockfiles-are-generally-available/)
- Agents: GitHub Agentic Workflows went from technical preview (Feb 13, 2026) to public preview (June 11, 2026) to GA (Oct 28, 2026) on Free, Team, and Enterprise; Markdown files in `.github/workflows/` compile to standard Actions workflows, and the default Copilot CLI engine needs a Copilot license — [Preview](https://github.blog/changelog/2026-02-13-github-agentic-workflows-are-now-in-technical-preview/); [Public preview](https://github.blog/changelog/2026-06-11-github-agentic-workflows-is-now-in-public-preview/); [GA](https://github.blog/changelog/2026-10-28-github-agentic-workflows-generally-available/)
- Agent costs: Actions minutes plus inference; 1 AI Credit = $0.01, with a default cap of 1,000 AIC per run — [GitHub Docs](https://docs.github.com/en/copilot/concepts/agent-capabilities/agent-powered-workflows-and-operations/about-github-agentic-workflows). Copilot coding agent and code review were among the services affected by the August 6 Actions incident — [August report](https://github.blog/news-insights/company-news/github-availability-report-august-2026/)
- AI in CI/CD generally: 48.8% of organizations use AI for at least one CI/CD task, 22% use agents in CI/CD, and only 8% run an AI step in the build — [JetBrains State of CI/CD 2026](https://www.jetbrains.com/pages/state-of-cicd2026/)

### Inferences
- Integration is the moat: a team already on GitHub gets CI, identity, policy, and agent execution with one vendor. The same concentration is why GitHub control-plane outages hit code, CI, and agents together.
- Rising adoption alongside a poor reliability record suggests buyers currently value integration and low operating burden over uptime. This is tentative: the JetBrains tool-adoption survey has no stated date, so it cannot be lined up against the incident timeline.

### Gaps
- No reliable current count of Marketplace actions was collected.
- No independent measurement of operational cost savings compared with self-managed CI was found.
- Real-world quality or ROI data for Agentic Workflows after GA is not yet available, as GA was five days before this research.
