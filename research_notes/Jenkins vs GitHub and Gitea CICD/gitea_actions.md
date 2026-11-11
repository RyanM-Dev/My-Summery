# Gitea Actions as a Self-Hosted CI/CD Option (state as of 2 November 2026)

Evidence note: release and feature facts come mostly from Gitea and Forgejo primary sources. Practitioner evidence is thin and largely anecdotal. Most fetched pages were read through a summarizing tool, so verify exact wording against the source before quoting it.

## What is the current Gitea release and how does Actions work?

### Takeaway
Current stable is Gitea 28.3.0 (1 Nov 2026), paired with Gitea Runner 5.x, the renamed `act_runner` and a hard fork of nektos/act. Gitea orchestrates and an independently deployed Go runner executes jobs in Docker, Kubernetes pods, or directly on the host. The runner has had five major versions since May 2026, so it is moving fast and upgrades carry real churn.

### Cited Findings

**Versions and cadence**
- Gitea 28.3.0 was released 1 Nov 2026 as the third maintenance release of the 28 series; the post strongly recommends upgrading immediately — [Gitea blog](https://blog.gitea.com/release-of-28.3.0/)
- Gitea 28.0.0 (30 Sep 2026) dropped the historical `1.` prefix, so it is 28.0.0 rather than 1.28.0 — [Gitea blog](https://blog.gitea.com/release-of-28.0.0/)
- Recent timeline: 1.26.0 on 18 Apr 2026 ([post](https://blog.gitea.com/release-of-1.26.0/)), 1.27.0 on 12 Jul 2026 ([post](https://blog.gitea.com/release-of-1.27.0/)), then 1.27.2 (14 Aug), 1.27.3 (29 Aug), 28.1.0 (6 Oct), 28.2.0 (18 Oct) — [Gitea blog index](https://blog.gitea.com/)
- Docs version selector already lists 29-dev — [Gitea docs](https://docs.gitea.com/usage/actions/comparison)
- Runner timeline: 1.0.0 on 5 May 2026 ([post](https://blog.gitea.com/release-of-runner-1.0.0/)); 3.0.0 blog dated 31 Jul 2026 but its changelog says 5 Aug ([post](https://blog.gitea.com/release-of-runner-3.0.0/)); 4.0.0 on 24 Sep ([post](https://blog.gitea.com/release-of-runner-4.0.0/)); 5.0.0 on 8 Oct ([post](https://blog.gitea.com/release-of-runner-5.0.0/))
- Ubuntu 26.04 runner images were announced 22 Oct 2026 — [Gitea blog index](https://blog.gitea.com/)

**Architecture and act lineage**
- Actions has been built in since Gitea 1.19. Gitea delegates jobs to a standalone Go runner that the docs call "a hard fork of nektos/act" — [Gitea docs overview](https://docs.gitea.com/usage/actions/overview)
- `act_runner` was renamed Gitea Runner at 1.0.0: repo `gitea/runner`, binary `gitea-runner`, image `gitea/runner`. The last `act_runner` line was v0.6.x. The announcement credits nektos/act but does not itself use the word "fork" — [Gitea blog](https://blog.gitea.com/release-of-runner-1.0.0/)
- Gitea describes Actions as "stable enough for production" and says it uses it across its own gitea.com repositories — [Gitea docs overview](https://docs.gitea.com/usage/actions/overview). Gitea itself is developed on GitHub using GitHub Actions, per the rival project — [Forgejo comparison](https://forgejo.org/compare-to-gitea/)

**Labels and execution modes**
- Labels map jobs to environments: `docker://<image>` (container), `kubernetes://<image>` (pod), and `host` (runner machine). Host jobs run as the runner's user with no isolation, and the docs limit them to trusted repositories — [Gitea runner docs](https://docs.gitea.com/usage/actions/act-runner)
- Requirements: Gitea 1.21+, Git 2.34.1+, and Docker or Podman for container jobs — [Gitea runner docs](https://docs.gitea.com/usage/actions/act-runner)
- A job runs on a runner that has all requested labels, using the environment of the first label. Custom labels no longer exist in Gitea; runners declare labels at startup — [Gitea FAQ](https://docs.gitea.com/usage/actions/faq/)
- Advanced features need current runners: job summaries, `continue-on-error`, and cancel-time cleanup steps in 1.27 require Runner 2.0.0 — [Gitea blog](https://blog.gitea.com/release-of-1.27.0/)

**Kubernetes options**
- Community path, added in Runner 4.1.0 — [Gitea blog](https://blog.gitea.com/release-of-runner-5.0.0/). Two modes — [Gitea Kubernetes docs](https://docs.gitea.com/runner/installation/kubernetes/):
  - Job pods: unprivileged, no Docker daemon, Kubernetes 1.30+, but "Docker actions and `docker://` steps are not supported". Services share the pod network, so two services cannot use the same port, and only `--health-*` service options work.
  - Docker-in-Docker: all workflow features work, but the daemon is privileged and a job could break out to the node. Rootless DinD manifests and a Helm chart (`gitea/helm-actions`) exist.
- Enterprise path: Actions Runner Controller, a Kubernetes operator with scale sets and ephemeral runners, requires an activated Gitea Enterprise above 23.8.0 — [Gitea Enterprise docs](https://docs.gitea.com/enterprise/features/actions-runner-controller/)

**Runner release churn (operationally relevant)**
- 3.0.0: cache v2 on by default; host-escape container options stripped unless privileged; one process per `.runner` file — [Gitea blog](https://blog.gitea.com/release-of-runner-3.0.0/)
- 4.0.0: added `builtin:checkout` and shared S3 cache; cache traffic rerouted through the runner; `valid_volumes` glob semantics changed — [Gitea blog](https://blog.gitea.com/release-of-runner-4.0.0/)
- 5.0.0: six breaking changes, including jobs without `runs-on` failing and OTLP metric renames; added mTLS client certificates — [Gitea blog](https://blog.gitea.com/release-of-runner-5.0.0/)
- 28.0.0: completed runs, logs, and artifacts are deleted after 400 days by default unless `RUN_RETENTION_DAYS = 0` is set before upgrading — [Gitea blog](https://blog.gitea.com/release-of-28.0.0/)

### Inferences
- The numbering change and fast major bumps mean version-specific claims age quickly. Anything written about "act_runner" or Gitea 1.2x before mid-2026 is likely stale.
- Runner upgrades should be treated as planned changes with staging tests, not background patching.

### Gaps
- I did not find a stated Gitea-to-runner version compatibility matrix beyond the Gitea 1.21 minimum.
- I did not fetch the 4.1.0 runner post directly; its Kubernetes backend date comes from the 5.0.0 post.

## What is compatible and incompatible with GitHub Actions?

### Takeaway
Core syntax is largely compatible, and 2026 releases closed major gaps (concurrency, first-class reusable workflows, artifacts v4 and cache via stock actions, dynamic matrices, token permissions). The main remaining gaps for serious teams are that `environment` is ignored, so there are no deployment protection rules, and OIDC is merged only for the unreleased 29.0.0.

### Cited Findings

| Feature | State as of 28.3.0 | Source |
|---|---|---|
| Overall | "Designed to be compatible", with listed differences; standard expression functions and contexts supported | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| Reusable workflows | First-class jobs since 1.27.0 (own entry, logs, graph node). **Breaking:** external reusable workflows by URL no longer supported; clone locally. 28.0: public repos cannot call private reusable workflows, and nested workflows cannot exceed caller token permissions | [1.27.0](https://blog.gitea.com/release-of-1.27.0/), [28.0.0](https://blog.gitea.com/release-of-28.0.0/) |
| Concurrency | Added in 1.26.0, flagged as breaking for overlapping pipelines. No `queue: max`; an issue reports queued work silently discarded with `cancel-in-progress: false`. A related cancel bug was fixed in 28.0.0 | [1.26.0](https://blog.gitea.com/release-of-1.26.0/), [#39195](https://github.com/go-gitea/gitea/issues/39195), [#39461](https://github.com/go-gitea/gitea/pull/39461) |
| OIDC | Not in 28.x: `id-token` is listed as unsupported. PR #39052 merged 18 Oct 2026, milestone 29.0.0: issuer at `/api/actions/oidc`, GitHub-compatible claims, `id-token: write` gating, none for fork PRs | [Comparison docs](https://docs.gitea.com/usage/actions/comparison), [PR #39052](https://github.com/go-gitea/gitea/pull/39052) |
| Environments | `jobs.<job_id>.environment` is ignored | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| Permissions | `permissions` controls `GITEA_TOKEN`, clamped by repo and owner settings. Unsupported: `statuses`, `checks`, `deployments`, `id-token`, `security-events`, `pages`. Gitea-only: `code`, `releases`, `wiki`, `projects` | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| Package publish | `GITEA_TOKEN` cannot publish to its repo's package registry, such as OCI push; a personal access token is suggested | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| Action resolution | Host-less `uses:` resolves to `https://github.com`, or the local instance when `[actions].DEFAULT_ACTIONS_URL = self`. Absolute URLs, `self:`, `builtin:`, and `$/` are Gitea extensions | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| Matrix | Dynamic matrices from job outputs, `strategy.max-parallel`, `needs`-dependent `runs-on`, and enforced `fail-fast` in 28.0. Job-level `if:` is now evaluated before matrix expansion with restricted contexts | [28.0.0](https://blog.gitea.com/release-of-28.0.0/) |
| Artifacts and cache | Since Runner 3.0.0, stock `upload-artifact`/`download-artifact` v4.4.0+ and `actions/cache` work without the old fork. Shared S3 cache in Runner 4.0.0 | [Runner 3.0.0](https://blog.gitea.com/release-of-runner-3.0.0/), [Runner 4.0.0](https://blog.gitea.com/release-of-runner-4.0.0/) |
| Services | Supported, but job logs have no separate service-step section. Job-pod mode has port and option limits | [Comparison docs](https://docs.gitea.com/usage/actions/comparison), [Kubernetes docs](https://docs.gitea.com/runner/installation/kubernetes/) |
| Ignored or missing | Problem matchers and error annotations ignored; pre/post steps have no separate log section; context availability is not checked | [Comparison docs](https://docs.gitea.com/usage/actions/comparison) |
| PR ref | For `pull_request`, `ref` is `refs/pull/:prNumber/head`, not a merge-preview ref | [Gitea FAQ](https://docs.gitea.com/usage/actions/faq/) |

- Outdated claim to discard: a migration-search summary reported a third-party source saying concurrency groups are ignored (exact source not pinned down), which conflicts with 1.26.0 — [Gitea blog](https://blog.gitea.com/release-of-1.26.0/)
- Workflows generally move across, but some actions hard-code GitHub-specific paths and may fail, and one guide advises pinning or using a compatible fork — [Elestio](https://blog.elest.io/gitea-actions-build-your-own-github-like-ci-cd-pipeline-without-the-price-tag/)

### Inferences
- Static-credential deploys to clouds remain necessary on any released Gitea until 29.0.0 ships. Teams should expect some early-version roughness since the feature merged only in October 2026.
- `ref` semantics differ from GitHub's merge ref, so PR CI tests the head, not the merge result, unless the workflow handles merging itself.
- Compatibility is real at the YAML level but not universal at the marketplace level. Actions calling GitHub-only APIs, such as `statuses` or `checks`, are the likely failures. No source quantifies the failure rate.

### Gaps
- I found no systematic compatibility test of marketplace actions against Gitea.
- `workflow_call` secrets inheritance details and limits were not verified.
- The comparison docs do not mention concurrency, reusable workflows, artifacts, or cache, so those rows rely on release notes.

## What gaps do users report?

### Takeaway
The main gaps are the absence of environment-gated deployments, autoscaling limited to Enterprise or DIY, a simple secrets model with no environment scope in the docs, and audit logs limited to Enterprise. Debugging UI improved materially in 2026 but still has known holes.

### Cited Findings
- **Auto-scaling:** the free tier documents no autoscaler; the Kubernetes page does not cover it, and the runner docs do not mention autoscaling — [Kubernetes docs](https://docs.gitea.com/runner/installation/kubernetes/), [Runner docs](https://docs.gitea.com/usage/actions/act-runner)
- "Kubernetes AutoScaling Runners" are listed as an Enterprise feature — [Gitea pricing](https://about.gitea.com/pricing/)
- A community KEDA scaler request exists; I could not confirm its outcome — [runner #456](https://gitea.com/gitea/runner/issues/456). Third-party `gitea-ci-autoscaler` provisions Hetzner nodes for K3s runners — [ecosyste.ms](https://awesome.ecosyste.ms/projects/github.com%2Frustunit%2Fgitea-ci-autoscaler)
- **Secrets scoping:** documented levels are user, organization, and repository; the more specific level wins. The page mentions no environment scope, global scope, or external vault integration — [Gitea secrets docs](https://docs.gitea.com/usage/actions/secrets/)
- **Required approvals:** fork PR approval exists, and gitea.com requires it. The docs do not say it is the default on every instance — [Gitea FAQ](https://docs.gitea.com/usage/actions/faq/). With `environment` ignored, deployment approval through GitHub-style environments is not available — [Comparison docs](https://docs.gitea.com/usage/actions/comparison)
- Manual-gate request: an open-looking issue asks for manually triggered jobs inside a run — [#38776](https://github.com/go-gitea/gitea/issues/38776) (status not verified)
- **Debugging UI:** added in 2026: dependency graph, re-run failed jobs, run summary (1.26) — [Gitea blog](https://blog.gitea.com/release-of-1.26.0/); job summaries (1.27) — [Gitea blog](https://blog.gitea.com/release-of-1.27.0/); build queue view, artifact preview, failed-run reporting for invalid workflows (28.0) — [Gitea blog](https://blog.gitea.com/release-of-28.0.0/)
- Remaining UI gaps: no error annotations or problem matchers, and no distinct log sections for service or pre/post steps — [Comparison docs](https://docs.gitea.com/usage/actions/comparison)
- **Audit:** audit log and SAML SSO are Enterprise features — [Gitea Enterprise](https://about.gitea.com/products/gitea-enterprise/)
- **Observability:** Runner 4.0.0 adds OpenTelemetry tracing; 5.0.0 adds `gitea_runner_state` and changes metric names — [Runner 4.0.0](https://blog.gitea.com/release-of-runner-4.0.0/), [Runner 5.0.0](https://blog.gitea.com/release-of-runner-5.0.0/)
- **Operational friction (anecdotal, seen only in search excerpts, pages not read in full):** a runner "kept getting wedged after upstream Gitea releases" — [lanakod](https://lanakod.com/en/projects/gitea-act-runner); difficulty using a private-registry job image — [jayground8](https://jayground8.github.io/tags/gitea); Docker socket permissions as a common failure — [Elestio](https://blog.elest.io/gitea-actions-build-your-own-github-like-ci-cd-pipeline-without-the-price-tag/); running the runner in a separate network namespace to contain a compromised container — [excipio via HN reader, direct fetch returned 403](https://hn.jakechampion.name/reader/https://excipio.tech/blog/improving-my-self-hosted-actions-runner-setup/)

### Inferences
- Regulated teams needing "protected deploy with named approvers" would have to build this outside Actions, for example through a protected branch plus required PR reviews, or an external CD tool. This is my reasoning rather than a documented pattern.

### Gaps
- I found no authoritative Gitea doc describing a replacement for environment protection rules, and no issue confirming roadmap plans for it.
- I found no user survey quantifying which gaps matter most.

## What is the security record?

### Takeaway
2026 has been a very heavy CVE year, including an actively exploited RCE in CISA KEV (CVE-2026-60004) and a recurring Actions class: fork-PR approval bypasses and over-broad task tokens. High disclosure volume partly reflects scrutiny, but operators should plan for fast patching, closed registration, and untrusted-code isolation.

### Cited Findings

**Exploited in the wild**
- CVE-2026-60004, code injection in the `diffpatch` endpoint: ordinary repo write access leads to shell commands as the Gitea OS user; open registration makes it reachable by any visitor who can sign up. Fixed in 1.27.1. KEV-listed 25 Aug 2026 with a 28 Aug federal deadline. A reported incident dropped a miner-like payload in about 11 seconds — [Help Net Security](https://www.helpnetsecurity.com/2026/08/26/gitea-cve-2026-60004-exploited-in-the-wild/)
- CVSS 9.8, affecting 1.17 and later before 1.27.1, per secondary reporting — [SOC Prime](https://socprime.com/blog/cve-2026-60004-critical-gitea-rce-exploited-to-deploy-miner-like-payloads/)
- Canadian Centre for Cyber Security alert AV26-845 also reports CVE-2026-20896 as exploited per open-source reporting; I did not determine what that CVE covers — [Cyber Centre](https://cyber.gc.ca/en/alerts-advisories/gitea-security-advisory-av26-845)

**Actions-specific CVEs (2026)**
- 1.27.0: CVE-2026-58416, fork PR Actions task reading a private repo — [Gitea blog](https://blog.gitea.com/release-of-1.27.0/)
- 1.27.0: CVE-2026-58443, CVSS 9.6, public-only token writing to private PR head branches and able to trigger Actions runs; requires a token of a user with branch write access — [vpncentral](https://vpncentral.com/critical-gitea-flaw-lets-public-only-tokens-write-to-private-repository-branches/)
- 1.27.3: CVE-2026-66877, concurrency group pass dispatching approval-pending fork PR jobs — [Gitea blog](https://blog.gitea.com/release-of-1.27.3/)
- 28.0.0: CVE-2026-104632, -104626, -103670, -94205 (fork approval bypass, cancelled-job revival, fork runs cancelling trusted runs) and -96580 (matrix memory exhaustion) — [Gitea blog](https://blog.gitea.com/release-of-28.0.0/)
- 28.3.0: CVE-2026-118386 (review-triggered workflows running unapproved from a user-chosen commit), CVE-2026-112887 (task tokens listing all private repos), CVE-2026-116474 (job token GPG key producing false "verified" signatures), CVE-2026-116289 (public-only tokens creating org runner registration tokens, exposing job secrets) — [Gitea blog](https://blog.gitea.com/release-of-28.3.0/)
- GHSA-32h8-5xqg-h466, "fork workflow gets secrets and write token via review commit_id", rated High — [GitHub advisories](https://github.com/go-gitea/gitea/security/advisories)
- Runner 3.0.0: CVE-2026-73802, host-escape container options on non-privileged runners now stripped — [Gitea blog](https://blog.gitea.com/release-of-runner-3.0.0/)
- CVE-2026-58426, Artifacts V4 API, CVSS 9.6, is reported by a low-authority source with internally inconsistent dates (fix in 1.26.4, patch 20 May vs disclosure 3 July); verify via GHSA-hg5r-vq93-9fv6 — [LocateRisk](https://locaterisk.com/en/?p=9086)

**Volume**
- Counting CVEs listed in release posts: 3 in 1.26.0 ([post](https://blog.gitea.com/release-of-1.26.0/)), 45 in 1.27.0 ([post](https://blog.gitea.com/release-of-1.27.0/)), 21 in 28.3.0 ([post](https://blog.gitea.com/release-of-28.3.0/)). 28.0.0 lists 20 security advisories ([post](https://blog.gitea.com/release-of-28.0.0/)). These are my counts of summarized lists, so they may be off slightly.
- The advisory listing paginates to 18 pages at 10 per page, implying roughly 170-180 published advisories in total — [GitHub advisories](https://github.com/go-gitea/gitea/security/advisories) (inferred; the page gives no total)
- 2025-era example: CVE-2026-25714 is an incomplete fix of CVE-2025-68941 for token-public-only checks, fixed in 1.26.2, per a secondary database — [Vulert](https://vulert.com/vuln-db/gitea-incomplete-cve-2025-68941-fix-user-orgs-missing-checktokenpubliconly-switch-case-logic-fl-)

**Runner isolation and third-party actions**
- Gitea docs warn against using an untrusted runner or providing a runner to anything untrusted. On a private company instance the risk is lower since colleagues are accountable. `GITEA_TOKEN` is scoped to the job's repository — [Docs overview](https://docs.gitea.com/usage/actions/overview), [Gitea FAQ](https://docs.gitea.com/usage/actions/faq/)
- Isolation ranking from the docs: `host` has none ([Runner docs](https://docs.gitea.com/usage/actions/act-runner)); DinD is privileged and escapable; job pods are unprivileged with no service account token by default ([Kubernetes docs](https://docs.gitea.com/runner/installation/kubernetes/))
- Forgejo's security guide, applicable by analogy due to shared lineage, warns about `container.privileged`, network modes, `valid_volumes` exposure, registries logged in on the host, and missing service-container and I/O limits — [Forgejo docs](https://forgejo.org/docs/latest/admin/actions/security/)
- Runner 3.0.0 masks secrets in JSON, base64, and URL-encoded forms — [Gitea blog](https://blog.gitea.com/release-of-runner-3.0.0/)

**Air-gapped use**
- Default host-less actions come from github.com, so a default install has a live dependency on GitHub for third-party code — [Comparison docs](https://docs.gitea.com/usage/actions/comparison)
- Options: `DEFAULT_ACTIONS_URL = self` plus local mirrors, or `self:owner/repo@ref` — [Gitea FAQ](https://docs.gitea.com/usage/actions/faq/). `builtin:checkout` needs no Node in the job image — [Gitea blog](https://blog.gitea.com/release-of-runner-4.0.0/)
- Maintainers have reportedly discussed deprecating `DEFAULT_ACTIONS_URL` in favor of full local URLs. I saw this only in a search excerpt of [runner #259](https://gitea.com/gitea/runner/issues/259) because the direct fetch hit a sign-in page; treat as unverified and possibly dated
- A forum thread points air-gapped users to the same setting — [Gitea forum](https://forum.gitea.com/t/on-premise-air-gapped-gitea/10119)
- Historical note: older config defaulted to gitea.com mirrors, so older guides showing `DEFAULT_ACTIONS_URL = https://gitea.com` are stale — [CubePath tutorial](https://cubepath.com/docs/development-tools/gitea-advanced-configuration-and-ci-cd)

### Inferences
- Internet-exposed instances with open registration carry much higher risk, as the KEV-listed flaw shows.
- The repeated fork-approval bypasses suggest public-contribution instances should not run fork PR workflows on runners with network or secret access, at least until this class settles.
- Pinning third-party actions by full SHA and mirroring them locally is a reasonable control, though this is my recommendation and the sources did not state it as Gitea guidance.

### Gaps
- I found no reliable compiled count of 2024-2025 Gitea CVEs. Search returned almost only 2026 items, so the earlier record is unassessed here, not clean.
- I found no independent audit report or comparative CVE rate against Forgejo, GitLab, or Jenkins.
- The share of 2026 volume attributable to AI-assisted or newly intensified bug hunting is unknown.

## What about Gitea Enterprise, Cloud, pricing, governance, and Forgejo?

### Takeaway
Gitea is open core under a for-profit company, with autoscaling runners, audit, and SAML in paid tiers. Forgejo is the non-profit, GPL hard fork. Its Actions stresses "familiarity" over compatibility, shipped OIDC first, but lags Gitea on `permissions`, `continue-on-error`, and a production Kubernetes backend. Direct migration from Gitea is no longer simple to assume.

### Cited Findings

**Gitea commercial and governance**
- Enterprise price shows "$9.5" and "$19" per user/month with an asterisk ("1-year commitment required with online payment"); the fetched page does not make clear which applies. Free tier includes Gitea Actions; Enterprise adds Kubernetes autoscaling runners, audit logs, SAML SSO, and priority support with SLA. A 30-day trial and discounts for non-profits and small companies are available — [Gitea pricing](https://about.gitea.com/pricing/)
- Enterprise also lists org-level branch protection inheritance, dependency scanning, IP allowlist, mandatory 2FA, a rolling release model, and a SOC 2 Type 2 badge. No customer names appear — [Gitea Enterprise](https://about.gitea.com/products/gitea-enterprise/)
- Gitea Cloud is described only as "Fully managed Gitea, ready in minutes", with no prices on the pricing page — [Gitea pricing](https://about.gitea.com/pricing/)
- Forgejo's characterization, which is a competitor's view: Gitea is Open Core and controlled by a for-profit through domain and trademark, requires copyright assignment even for MIT code, and gives advance security notice to customers only — [Forgejo comparison](https://forgejo.org/compare-to-gitea/)

**Forgejo context**
- Created October 2022 after a for-profit took over Gitea; hard fork in early 2024; sits under non-profit Codeberg e.V. — [Forgejo comparison](https://forgejo.org/compare-to-gitea/)
- Forgejo v17.0 was released 15 Oct 2026 (non-LTS, supported to 28 Jan 2027). LTS 15.0 runs to 15 Jul 2027, and next LTS 19.0 is planned for 15 Apr 2027 — [Forgejo blog](https://forgejo.org/2026-10-release-v17-0/)
- Forgejo publishes a dated LTS schedule; I found no equivalent Gitea LTS commitment in the sources read.
- Security process: Forgejo offers advance notice to policy signatories, available to everyone — [Forgejo comparison](https://forgejo.org/compare-to-gitea/)

**Forgejo Actions vs Gitea Actions**

| Topic | Forgejo | Gitea | Source |
|---|---|---|---|
| Philosophy | "Strives for familiarity instead of compatibility" | "Designed to be compatible" | [Forgejo](https://forgejo.org/docs/latest/user/actions/github-actions/), [Gitea](https://docs.gitea.com/usage/actions/comparison) |
| OIDC | Available via `enable-openid-connect`, not `permissions: id-token` | Unreleased 29.0.0, using `id-token: write` | [Forgejo](https://forgejo.org/docs/latest/user/actions/github-actions/), [PR #39052](https://github.com/go-gitea/gitea/pull/39052) |
| Job `permissions` | Ignored | Controls `GITEA_TOKEN` | [Forgejo](https://forgejo.org/docs/latest/user/actions/github-actions/), [Gitea](https://docs.gitea.com/usage/actions/comparison) |
| `continue-on-error` at job level | Ignored | Supported since 1.27 | [Forgejo](https://forgejo.org/docs/latest/user/actions/github-actions/), [Gitea](https://blog.gitea.com/release-of-1.27.0/) |
| Default actions source | Example `https://data.forgejo.org`, curated Free Software | `https://github.com` | [Forgejo](https://forgejo.org/docs/latest/admin/actions/), [Gitea](https://docs.gitea.com/usage/actions/comparison) |
| Kubernetes job backend | Prototype as of May 2026; no production backend found | Runner 4.1.0+ | [Forgejo report](https://forgejo.org/2026-05-monthly-report/), [Gitea](https://blog.gitea.com/release-of-runner-5.0.0/) |
| Ephemeral runners | Documented in security guide | Exists (restart bug fixed in 5.0.0); Enterprise ARC adds scale sets | [Forgejo](https://forgejo.org/docs/latest/admin/actions/security/), [Gitea](https://blog.gitea.com/release-of-runner-5.0.0/) |
| Job summaries | v17 with Runner 13.2+ | 1.27 with Runner 2.0.0+ | [Forgejo](https://forgejo.org/2026-10-release-v17-0/), [Gitea](https://blog.gitea.com/release-of-1.27.0/) |
| Runner | Forgejo Runner, v13.0.0 on 3 Aug 2026 | Gitea Runner 5.x | [Forgejo news](https://forgejo.org/news/), [Gitea](https://blog.gitea.com/release-of-runner-5.0.0/) |

- Blogs conflict on who leads on compatibility, and one claims both use the same `act_runner`, which is wrong for current releases because each project ships its own runner — [codey.ch](https://www.codey.ch/forgejo-vs-gitea/), [techfuelhq](https://techfuelhq.com/homelab/forgejo-vs-gitea-2026/), [contabo](https://contabo.com/blog/gitea-vs-forgejo/)
- Migration from Gitea is described as possible but the page lists no supported source versions — [Forgejo comparison](https://forgejo.org/compare-to-gitea/)

### Inferences
- For teams wanting GitHub Actions portability and Kubernetes job pods, Gitea currently looks stronger. For teams prioritizing free-software governance, OIDC today, or LTS predictability, Forgejo looks stronger.
- Two years of divergence probably makes an in-place Gitea 28 to Forgejo upgrade unreliable, so migration should be treated as needing a test or a repository-level migration. Confirm against Forgejo's upgrade guide.

### Gaps
- Unable to confirm exact current Enterprise price or any Gitea Cloud price; check the pricing page directly.
- Could not confirm which Forgejo Actions PRs (expanded reusable workflows #10448, concurrency #9434, OIDC federation #10481) shipped in which release.
- No verified statement of the Gitea Ltd corporate structure, funding, or ownership was collected beyond Forgejo's description.

## What do practitioners report, who uses it at scale, and what does it cost?

### Takeaway
Evidence favors homelab and small-team success, with low server overhead and CI costs driven by the actual workloads. I found no public, credible case study of large-scale or regulated use, so claims of enterprise scale are vendor statements only.

### Cited Findings
- A January 2026 forum thread, low authority and pseudonymous: the poster moved CI to a home-server runner with existing YAML unchanged. One team reported over a year of small-team production use, needing a large EC2 instance for intensive testing; another reported about $4/month deploying Terraform to personal Azure; another reported smooth use since launch — [Yo Mother Board](https://yomotherboard.com/question/has-anyone-tried-running-gitea-actions-for-production-workloads/)
- Migration motivation: one guide ties interest to GitHub's December 2025 announcement of a $0.002/minute fee on self-hosted runners for private repos — [Elestio](https://blog.elest.io/how-to-migrate-from-github-actions-to-gitea-actions/). Verify GitHub's current position separately, since this is a secondary source.
- Resource claims: about 200 MB total for Gitea vs 8 GB minimum for GitLab, and about 50 MB idle in another table, both from low-quality sources without methodology — [Luca Berton](https://lucaberton.com/blog/gitea-self-hosted-git-server-kubernetes-2026/), [CloudInGenium](https://kx.cloudingenium.com/gitea-self-hosted-git-server-setup-migration-troubleshooting)
- Anecdote: a developer left GitLab citing RAM use, without measurements — [DEV Community](https://dev.to/genodev/no-more-gitlab-a-production-ready-gitea-stack-with-traefik-act-runner-and-auto-deploy-via-ssh-1g8e)
- A Habr article discusses runner memory, but numbers could not be extracted — [Habr](https://habr.com/en/articles/770448)
- Vendor claim: the runner "can scale to meet your needs for both small teams and large enterprises", no figures given — [Gitea Runner product page](https://about.gitea.com/products/runner/)
- Gitea's own dogfooding is the only named usage found — [Docs overview](https://docs.gitea.com/usage/actions/overview)
- Ongoing operating burden is visible in release notes: three runner major versions in about 10 weeks (Jul-Oct 2026), an emergency upgrade cycle for 1.27.1, and repeated "upgrade as soon as possible" guidance — [Runner 3.0.0](https://blog.gitea.com/release-of-runner-3.0.0/), [Runner 5.0.0](https://blog.gitea.com/release-of-runner-5.0.0/), [Help Net Security](https://www.helpnetsecurity.com/2026/08/26/gitea-cve-2026-60004-exploited-in-the-wild/), [28.3.0](https://blog.gitea.com/release-of-28.3.0/)

### Inferences
- The server is cheap, while runner capacity tracks your builds much like any self-hosted CI. The real cost is engineer time on patching, runner images, and caching.
- Default runner images lack GitHub-hosted images' tool breadth, as Forgejo documents for its own slim default, so teams should expect to build and maintain images. Gitea's Ubuntu runner images reduce but may not remove this.
- Suitable today for small and mid-sized teams on private instances. For large or regulated teams, the lack of environment approvals and released OIDC, plus an unproven record at scale, are the main objections.

### Gaps
- I found no named company case studies, published job throughput, concurrent-runner counts, or measured runner overhead.
- I found no regulated-industry report or audit experience with Gitea Actions. The SOC 2 badge covers the vendor, not a customer's deployment.
- Reddit, Hacker News, and conference talks were not systematically reviewed, so stronger practitioner reports may exist.
