# Jenkins, GitHub Actions, and Gitea Actions: Adoption, Migration, Cost, and Decision Criteria (as of 2 November 2026)

Evidence labels used throughout: **[Survey]** = published survey data; **[Vendor]** = claim by a party selling a product in this market; **[Case study]** = first-party named company or academic account; **[Anecdote]** = individual blog, forum, or press-reported opinion. "Secondary" means I could not read the original and relied on another source's report of it.

**Do-not-cite warning for the report writer:** several DEV Community posts by one author (`johalputt`) surfaced repeatedly with striking figures: "Netflix migrated 12,400+ pipelines and cut CI costs 35%", "ditched GitHub Actions for Tekton, cut costs 40%", and CI spend rising "$62k to $187k". They refer to products I could not confirm exist as described ("GitHub Actions 3.0", "Dependabot 2.0"), and I found no corroboration from Netflix or GitHub. They are excluded from the findings below as unverifiable. Example: [the "Netflix" post](https://dev.to/johalputt/case-study-netflix-migrated-from-jenkins-to-github-actions-2026-and-cut-ci-costs-35-b49).

## What do surveys say about adoption and trends?

### Takeaway
GitHub Actions leads every recent survey that reports tool share, with its lead widening in JetBrains' 2026 data (47% of organizations, up from 41%), while Jenkins is declining slowly but remains in roughly a quarter to two-fifths of samples. Gitea Actions and Forgejo Actions do not appear in any survey data I found, so their adoption is unmeasured rather than demonstrably small.

### Cited Findings

**JetBrains State of CI/CD 2026 [Survey; publisher sells TeamCity]**

- Datasets: CI/CD Tools Adoption Survey 2026 with 598 respondents, of whom 576 use at least one CI/CD tool; Developer Ecosystem Survey 2026 with 21,946 respondents in the page's intro but 22,082 in its methodology section, fielded May to July 2026. Recruitment was via developer blogs, newsletters, social media, and community lists, weighted, and the page says it reflects engaged developers rather than the whole industry — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- Organization-level tool use, 2026 vs 2025: GitHub Actions 47 vs 41; GitLab CI 32 vs 34; Jenkins 28 vs 31; Azure DevOps Server 15 vs 14; TeamCity 10 vs 7; Bitbucket Pipelines 6 vs 7; Google Cloud Build 4 vs 4 — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- JetBrains' two 2026 surveys disagree with each other: Jenkins was 28% in one and 37% in the other, and TeamCity 9% vs 4%. JetBrains attributes the gap to organization vs individual framing and different recruitment. The page does not label which is which; 28% matches the organization chart — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- Hosting of the primary CI/CD tool: 72.8% self-host, composed of 30.5% on own premises (described as a five-year high) and 42% on self-managed cloud instances; 24.4% use a fully managed service, down seven points from a 2022 peak — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- 83% run CI/CD pipelines (82% in 2025); GitOps 23% (29% in 2025); platform engineering 16% (17%) — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- Multi-tool use: 25% use two or more tools and 17% use three or more. These read oddly together and the page's comparison baselines (33% two-tool, 12% three-plus) conflict with the 2025 figures below, so treat them as loose — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- AI: about 49% of organizations use AI for at least one CI/CD task, 22% report AI agents in CI/CD, and only 8% run AI steps inside builds — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)

**JetBrains State of CI/CD 2025 [Survey; secondary for sample]**

- 805 respondents; 62% of individuals and 41% of organizations reported GitHub Actions; 32% of organizations used two tools and 9% at least three. The original returned HTTP 403, so sample and individual figures rest on a search summary of it — [JetBrains blog, Oct 2025](https://blog.jetbrains.com/teamcity/2025/10/the-state-of-cicd/). The organization figures are confirmed by the 2026 page's comparison column — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- A separate Japanese report headlines 41% of organizations using multiple CI/CD tools (32% + 9%) — [@IT, Oct 2025](https://atmarkit.itmedia.co.jp/ait/articles/2510/14/news037.html)

**JetBrains State of Developer Ecosystem 2025 [Survey; secondary]**

- Reported organizational use of GitHub Actions 33%, Jenkins 28%, GitLab CI 19%, and 55% of developers regularly using CI/CD tools. This came from a search summary of a JetBrains article that returned 403 when fetched, so it is unconfirmed — [JetBrains blog, Mar 2026](https://blog.jetbrains.com/teamcity/2026/03/best-ci-tools)

**CNCF Annual Survey 2024, published April 2025 [Survey; primary PDF read]**

- Sample: 689 for the tool questions (988 in 2023); the report's prose says 750 members responded in fall 2024 — [CNCF/Linux Foundation PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en)
- "Tools in use to manage CI/CD pipelines", select all, 2024 vs 2023: GitHub Actions 51 vs 43; Argo 45 vs 39; Jenkins 44 vs 32; GitLab 43 vs 36; Tekton Pipelines 5 vs 5. I read these from extracted PDF chart text and cross-checked them against the printed growth callouts (+19%, +16%, +40%, +20%); the Jenkins callout is slightly above the 37.5% those bars imply — [CNCF/Linux Foundation PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en)
- CI/CD used in production for most or all applications rose from 46% (2023) to 60% (2024) — [CNCF/Linux Foundation PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en)
- **Correction of widely repeated secondary numbers:** press coverage gives GitLab as 34% — [Cloud Native Now](https://cloudnativenow.com/topics/cloudnativedevelopment/cncf-survey-surfaces-steady-pace-of-increased-cloud-native-technology-adoption/). The primary chart shows 43%. A "Best of 2025" reprint also makes these 2024-survey numbers look like 2025 data — [Cloud Native Now](https://cloudnativenow.com/editorial-calendar/best-of-2025/cncf-survey-surfaces-steady-pace-of-increased-cloud-native-technology-adoption-2/)

**CNCF Annual Survey 2025, published 20 January 2026 [Survey; CI/CD data not obtained]**

- Headline: 82% of container users run Kubernetes in production — [CNCF](https://www.cncf.io/reports/the-cncf-annual-survey-2025). Sample is disputed: 628 per [TFIR](https://tfir.io/cncf-survey-cloud-native-maturity-2025/); 689 per [FAUN](https://faun.dev/c/news/devopslinks/the-cloud-native-tipping-point-what-689-companies-just-revealed), which may be confusion with the 2024 survey.

**CNCF Technology Radar Q1 2026 [Survey; perception, not share; primary PDF read]**

- More than 400 professional cloud-native developers overall, n=363 for workflow automation and n=352 on rating charts. GitHub Actions, Armada, Buildpacks, Jenkins, and Argo CD were placed in "adopt"; Karmada, Tekton, Jenkins X, and werf in "trial" — [CNCF Radar](https://www.cncf.io/wp-content/uploads/2026/03/Q1-2026-CNCF-Technology-Radar-Report.pdf)
- Maturity rated 4 or 5 stars: GitHub Actions 89% (45% 5-star), Jenkins 85% (47% 5-star, the highest 5-star share in the category), Tekton 76% — [CNCF Radar](https://www.cncf.io/wp-content/uploads/2026/03/Q1-2026-CNCF-Technology-Radar-Report.pdf)
- Likely or highly likely to recommend: GitHub Actions 91%, Argo CD 87%, Tekton 85%, Jenkins 66%, placing Jenkins near the bottom of the list despite its maturity score — [CNCF Radar](https://www.cncf.io/wp-content/uploads/2026/03/Q1-2026-CNCF-Technology-Radar-Report.pdf)

**GitHub Octoverse 2025 [Vendor telemetry; data 1 Sep 2024 to 31 Aug 2025]**

- 180M+ developers on GitHub; 11.5B GitHub Actions CPU minutes in public projects, up 35% from 8.5B in 2024; 13.5B including self-hosted under the prior year's method, up 30% — [GitHub](https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/). This covers public projects only and says nothing about market share, though 81.5% of contributions were in private repositories.

**Stack Overflow Developer Survey 2025 [Survey; no CI tool data visible]**

- The technology page's visible content has no CI tool table: Docker 71.1% and Kubernetes 28.5% (n=24,473 for cloud development tools); GitHub 81.1% and GitLab 35.6% for code documentation and collaboration (n=30,065). The retrieved page was truncated, so a CI table elsewhere cannot be fully excluded — [Stack Overflow](https://survey.stackoverflow.co/2025/technology)

**CD Foundation State of CI/CD [Survey; latest edition found is 2024]**

- I found no 2025 or 2026 edition. The 2024 report, seen via a search summary and not read in full, found that growing use of self-hosted CI/CD tools did not necessarily lead to greater performance, possibly because of interoperability issues — [CD Foundation, Apr 2024](https://cd.foundation/blog/2024/04/16/state-cicd-devops-tooling-adoption/)

**DORA 2025 [Survey; AI focus]**

- Nearly 5,000 technology professionals and 100+ hours of qualitative data. Results retrieved covered AI adoption (90%), with no CI tool ranking — [Google](https://blog.google/innovation-and-ai/technology/developers-tools/dora-report-2025/); [report page](https://cloud.google.com/devops/state-of-devops)

### Inferences
- Jenkins' CNCF jump (32% to 44%) occurred alongside a sample drop from 988 to 689, and JetBrains shows Jenkins declining through 2025 to 2026. I would not read the CNCF move as a revival. Sample composition is a likely contributor, not a proven cause.
- The gap between surveys (Jenkins 28%, 37%, 44%) mostly shows sampling differences. A defensible statement is "GitHub Actions first, Jenkins and GitLab CI competing for second, with order depending on sample."
- JetBrains' 72.8% self-hosted and 47% GitHub Actions are not contradictory on their face, because they come from different question sets and "self-managed cloud instances" counts as self-hosting. But the page gives no cross-tab, so this cannot be taken further.
- Jenkins' pattern in the CNCF radar, high maturity and low recommendation, fits "trusted but not chosen again". This is my interpretation of one perception survey.

### Gaps
- **No survey data on Gitea Actions or Forgejo Actions.** The JetBrains 2026 page does not mention them, and neither appears in the CNCF 2024 top 15. No reliable adoption figure exists in what I found.
- CNCF 2025 report CI/CD figures: the PDF exceeded the fetch tool's size limit and was not read. The 2024 figures are the latest CNCF tool shares I can vouch for.
- No company-size or regional breakdown of tool share was available from the JetBrains 2026 page.
- The 2025 JetBrains blog and March 2026 article were blocked (403), so their figures above are not verified against the originals.

## What documented migrations exist, what did they cost, and what are the Importer's limits?

### Takeaway
Credible first-party accounts cite operational relief and faster pipelines, but none I found gives a full migration cost in person-months, and the best measured result is explicitly confounded by refactoring. GitHub Actions Importer publishes no conversion rate and leaves scripted pipelines, secrets, runners, and unknown plugins to manual work.

### Cited Findings

**Jenkins to GitHub Actions**

- **Bankdata (Denmark, banking IT), arXiv 2026 [Case study; academic, authors include migration team]:** moved a COBOL mainframe pipeline with Groovy across four repos. Mean runtime fell from 724 s (SD 47.3) to 130 s (SD 9.5), an 82.1% reduction across five runs per configuration; container image shrank 1,423 to 559 MiB — [arXiv 2604.00936](https://arxiv.org/html/2604.00936v1)
    - What got worse: less step-level visibility, so failures were harder to localize; COBOL developers found the move to GitHub's web interface disruptive; masking secrets in logs was judged insufficient alone — [arXiv 2604.00936](https://arxiv.org/html/2604.00936v1)
    - Caveats: the authors call the speedup an upper bound because it mixes architecture and platform effects; a controlled refactored-on-Jenkins run was blocked; one pipeline only; no deployment testing — [arXiv 2604.00936](https://arxiv.org/html/2604.00936v1)
- **FreeAgent, June 2022 (older) [Case study]:** Rails monolith with 58,000+ tests. Slowest deployments' p95 fell from 25 to 16 minutes and monthly test-infrastructure spend fell 90%, using Actions with ephemeral self-managed EC2 Fleets — [FreeAgent Engineering](http://engineering.freeagent.com/2022/06/10/in-full-flow-moving-from-jenkins-to-actions-part-1/)
    - Reasons to leave: single controller limits halting all builds, poisoned agents, hard-to-test upgrades, EC2 plugin limits on instance types, unreviewed settings, and Groovy outside team expertise. Migration effort is not stated in Part 1 — [FreeAgent Engineering](http://engineering.freeagent.com/2022/06/10/in-full-flow-moving-from-jenkins-to-actions-part-1/)
- **Karthik Gumpeni, June 2025 [Anecdote]:** 85+ pipelines, three Jenkins servers retired, about 30% lower CI time with no baseline given, zero production incidents, over "a few months". No downsides reported — [DEV](https://dev.to/karthik_gumpeni/from-jenkins-to-github-actions-my-cicd-migration-journey-30p1)
- **Consultancy guidance [Vendor-adjacent]:** think in pipeline archetypes rather than line-by-line Jenkinsfile conversion, since events work differently — [CBTW](https://cbtw.tech/insights/lessons-learned-when-migrating-from-jenkins-to-github-actions)

**GitHub Actions Importer coverage and limits [Vendor documentation]**

- Fully supported: Docker templates in freestyle jobs; `docker` and `stage` in Jenkinsfile. Partial: `agent`, `environment`, `stages`, `steps`, `triggers`, `when`. **Not supported:** `inputs`, `matrix`, `options`, `parameters` — [GitHub Docs](https://docs.github.com/en/actions/tutorials/migrate-to-github-actions/automated-migrations/jenkins-migration)
- Manual tasks: mandatory build tools, **scripted pipelines**, secrets, self-hosted runners, unknown plugins, unsupported build steps — [GitHub Docs](https://docs.github.com/en/actions/tutorials/migrate-to-github-actions/automated-migrations/jenkins-migration)
- No overall conversion rate is published; each pipeline is classified as successful, partially successful, unsupported, or failed only after running `audit` on your own instance — [GitHub Docs](https://docs.github.com/en/actions/tutorials/migrate-to-github-actions/automated-migrations/jenkins-migration)
- The plugin table has about 85 entries, and about 17 listed as having no Actions equivalent, including Lock, Retry, Mailer, TimeStamper, Warnings. Unlisted plugins are left as comments in the output. These counts are approximate, from an automated read of the page — [gh-actions-importer docs](https://raw.githubusercontent.com/github/gh-actions-importer/main/docs/jenkins/index.md)

**Moves off GitHub or to Gitea/Forgejo**

- **Dutch government code.overheid.nl, soft launch 24 April 2026 [Secondary press]:** self-hosted Forgejo for agencies; early adopters reported as four ministries and four cities; reasons given as sovereignty, data control, governance, and jurisdiction — [byteiota](https://byteiota.com/netherlands-builds-github-rival-forgejo-migration-2026/). No cost or CI outcome data published in what I read.
- **Zig, Nov 2025 [Anecdote; press-reported]:** moved to Codeberg after about ten years, citing Actions unreliability (unpredictable scheduled jobs, an uncleared CI backlog) and GitHub's AI focus — [Techzine](https://www.techzine.eu/news/devops/136914/zig-project-leaves-github-due-to-excessive-ai/). A safe-sleep bug went unaddressed for months, and Kelly later apologized for his post's tone — [The Register](https://www.theregister.com/2025/12/02/zig_quits_github_microsoft_ai_obsession/?td=rt-3a). Cost: GitHub Sponsors was the foundation's largest 2024 income source — [DevClass](https://www.devclass.com/ci-cd/2025/11/27/zig-project-ditches-github-for-codeberg-but-move-could-be-costly/1727053)
- **Jorijn, May 2026 [Anecdote; summary only]:** individual move to Forgejo v15 LTS on one NUC with Postgres and Traefik, framed as sovereignty rather than reliability. Trade-offs: Actions ecosystem friction, no enterprise support, Renovate replacing Dependabot. His original page returned unrelated content when fetched — [daily.dev summary](https://daily.dev/posts/why-i-m-leaving-github-for-forgejo-8ezk0te27.md)

**Compatibility, the main migration cost for Gitea or Forgejo [Vendor documentation]**

- Gitea: `jobs.<job_id>.environment` ignored; problem matchers and error annotations ignored; `GITEA_TOKEN` cannot push to its own repository's package registry (PAT workaround); `id-token` and other GitHub-only permission scopes unsupported; unqualified actions download from github.com by default unless `DEFAULT_ACTIONS_URL` is set to `self` — [Gitea Docs](https://docs.gitea.com/usage/actions/comparison)
- Forgejo: states it aims for familiarity and "is not designed to be compatible". Runners default to Debian bookworm with only Node.js; `permissions` and `continue-on-error` subkeys are ignored; OIDC uses `enable-openid-connect` rather than `id-token: write`; workflows can also run in LXC — [Forgejo Docs](https://forgejo.org/docs/latest/user/actions/github-actions/)
- One secondary source reports `actions/checkout@v6` broke authenticated checkout on non-GitHub runners in early 2026; I could not confirm this against official documentation — [byteiota](https://byteiota.com/netherlands-builds-github-rival-forgejo-migration-2026/)

### Inferences
- Importer's unsupported list (`matrix`, `parameters`, scripted pipelines, unknown plugins) overlaps with what makes mature Jenkins estates hard, so its audit is best used as a scoping tool rather than a conversion engine. The docs do not state that framing; it is my reading.
- FreeAgent and Bankdata both changed architecture while migrating (ephemeral runners, containers), so neither isolates platform effect. Treat "Actions made it faster" claims as "the rewrite made it faster."
- Gitea's default github.com action resolution means an air-gapped or restricted site must mirror actions and set `DEFAULT_ACTIONS_URL`. Inferred from the Gitea doc.
- Keyless cloud deployments using `id-token` will need rework when moving to Gitea, and syntax changes on Forgejo. Inferred from the two compat docs, which also differ from each other, so Gitea and Forgejo should not be treated as interchangeable.

### Gaps
- **No verified person-month or dollar cost for a large Jenkins migration** from a first-party source. FreeAgent Part 2 might have approach details, but was not read.
- No published Jenkins to Gitea or Jenkins to Forgejo case study with metrics found.
- No first-party engineering account of leaving GitHub Actions for self-hosted systems with metrics; results were vendor pages or the excluded posts.
- Importer conversion rates at real estates: none published.

## Is there published total-cost evidence?

### Takeaway
I found no independent, assumption-stated TCO study comparing all three. What exists is firm list pricing, GitHub's own impact claims, and a few worked examples, with the crucial cost for self-hosted Jenkins and Gitea (engineer time) unmeasured.

### Cited Findings
- **Hosted runners cut, in effect 1 Jan 2026:** up to 39% depending on machine type — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- **$0.002/min self-hosted platform fee announced Dec 2025 for 1 Mar 2026, then postponed** with no new date while GitHub re-evaluates — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/)
- **Current state:** GitHub's billing docs say self-hosted usage "remains free" and does not reduce included usage. Linux 2-core x64 is $0.006/min, arm64 $0.005/min. Included allowance is expressed in dollars: Free $12, Pro and Team $18, Enterprise Cloud $300 — [GitHub Docs](https://docs.github.com/en/billing/concepts/product-billing/github-actions) (read 2 Nov 2026)
- GitHub's impact claim for the planned change: 96% of customers no change; of the 4% affected, 85% lower bills and 15% a median increase of about $13 — [GitHub Changelog](https://github.blog/changelog/2025-12-16-coming-soon-simpler-pricing-and-a-better-experience-for-github-actions/) **[Vendor]**
- Press worked example: 20 developers × 10 builds/day × 8 min × 22 days = 35,200 min, about $70/month at the postponed fee. A Reddit user with 1.6M monthly minutes expected a $3,500 rise (arithmetic gives $3,200) — [The Stack](https://www.thestack.technology/github-to-charge-self-hosted-actions-users-what-does-that-look-like/) **[Anecdote for the Reddit figure]**
- Atlassian announced a similar self-hosted runner charge for Bitbucket Pipelines weeks earlier — [The Stack](https://www.thestack.technology/github-to-charge-self-hosted-actions-users-what-does-that-look-like/)
- Switching drivers: among those planning to switch, scaling costs 24%, maintenance effort 23%, consolidation 22%, complex configuration 21%. The base is not stated; 62% have no plans to switch — [JetBrains 2026](https://www.jetbrains.com/pages/state-of-cicd2026/) **[Survey]**
- FreeAgent's 90% test-infrastructure saving (2022) came from self-managed ephemeral EC2 fleets, not from hosted minutes — [FreeAgent Engineering](http://engineering.freeagent.com/2022/06/10/in-full-flow-moving-from-jenkins-to-actions-part-1/) **[Case study]**
- One team says server cost was not the real expense, with maintenance the larger hidden cost (no measured figure, from a search summary) — [Medium, Oct 2025](https://bhavyansh001.medium.com/why-we-replaced-jenkins-with-github-actions-and-never-looked-back-c6b0eb9e5ea3) **[Anecdote]**

### Inferences
- **Illustrative arithmetic only, my own, list price, Linux 2-core x64, before allowance:** using The Stack's 35,200 min/month, hosted is 35,200 × $0.006 = **$211.20/month**; self-hosted platform fee today is **$0** plus your own compute and labour. At 1.6M min/month, hosted is **$9,600/month**. Excludes storage, larger runners, and egress. It does not include any engineer cost, because I found no sourced figure.
- For small and moderate volumes, the licence or minutes line is small compared with one engineer's time, so maintenance load is likely to decide the outcome. This is reasoning supported by the switching-reason data, not a measured result.
- GitHub's postponement leaves pricing risk open: budgets should test the scenario in which $0.002/min returns, as Actions orchestration is free today by policy rather than by contract as far as I can see.

### Gaps
- **No independent TCO study for Jenkins vs GitHub Actions vs Gitea Actions.** I would treat any precise comparative percentage as unverified.
- No sourced staffing figures for operating Jenkins or Gitea at a given scale.
- No published Gitea Actions at-scale cost data. The OpenMetal cost article could not be read, so it is not used.

## What criteria do senior engineers actually use, and is data sovereignty significant in 2025-2026?

### Takeaway
No survey measures senior engineers' decision criteria directly; the evidence comes from switching reasons and case studies, and it points to maintenance, cost at scale, consolidation, and SCM location. Sovereignty is a real and rising driver in European public-sector and individual moves, with concrete evidence in the Netherlands, while sanctions restrictions are narrow and specific, not general.

### Cited Findings

**Operational and team**

- Switching reasons (see above): scaling costs, maintenance effort, consolidation, complex configuration — [JetBrains 2026](https://www.jetbrains.com/pages/state-of-cicd2026/) **[Survey]**
- SCM gravity: FreeAgent already paid for GitHub and ran static checks on Actions, and YAML was easier than Groovy for a Ruby team — [FreeAgent](http://engineering.freeagent.com/2022/06/10/in-full-flow-moving-from-jenkins-to-actions-part-1/) **[Case study]**
- Lock-in and special hardware: Bankdata's goals included less vendor lock-in and portability, with mainframe compatibility as a constraint. It kept a monolithic script for portability, accepting worse step visibility, and portability to other CI remained untested — [arXiv](https://arxiv.org/html/2604.00936v1) **[Case study]**
- Reliability: Zig's reasons were led by Actions behaviour — [Techzine](https://www.techzine.eu/news/devops/136914/zig-project-leaves-github-due-to-excessive-ai/) **[Anecdote]**. One outlet reports 257 GitHub incidents and 48 major outages from May 2025 to April 2026; I did not verify this against GitHub's status history — [byteiota](https://byteiota.com/netherlands-builds-github-rival-forgejo-migration-2026/)

**Sovereignty in Europe**

- The Dutch government launched Forgejo-based code.overheid.nl on 24 April 2026, citing sovereignty, data control, and governance, with FISA 702 and the CLOUD Act named as concerns — [byteiota](https://byteiota.com/netherlands-builds-github-rival-forgejo-migration-2026/) (secondary)
- A government news item dated 2 October 2026, "New push for open source in government and education", reports national government, provinces, municipalities, and SURF agreeing to make open source a more systematic part of operations — [NL Digital Government](https://www.nldigitalgovernment.nl/news/)
- Individual engineers cite FISA 702 and the CLOUD Act as unresolvable, along with GitHub's absorption into Microsoft CoreAI and an April 2026 opt-out Copilot training change — [daily.dev summary](https://daily.dev/posts/why-i-m-leaving-github-for-forgejo-8ezk0te27.md) **[Anecdote]**
- Codeberg's supporting membership reportedly about doubled, from 600 to over 1,200, around the Zig move — [The Register](https://www.theregister.com/2025/12/02/zig_quits_github_microsoft_ai_obsession/?td=rt-3a). Small absolute numbers.
- **Vendor response:** GitHub Enterprise Cloud EU data residency, generally available from 29 Oct 2024 on ghe.com, with Zeiss and CARIAD as early adopters — [DevClass, Sep 2024](https://devclass.com/2024/09/25/github-adds-microsoft-azure-based-eu-data-residency-to-enterprise-cloud). Data may be replicated across regions for resilience and AI services may process elsewhere, per the same analysis.
- **GitHub Enterprise Local** (GHES on Azure Local) is described as preview, for disconnected operation, with Actions on self-hosted runners; the source is an undated third-party post — [schneider.im](https://www.schneider.im/github-enterprise-local-devops-for-sovereign-and-private-cloud-environments/)
- Survey signal is indirect: among organizations not using AI in CI/CD, 36% cite privacy, security, or IP concerns and 14% regulatory restrictions. This concerns AI, not CI hosting — [JetBrains 2026](https://www.jetbrains.com/pages/state-of-cicd2026/) **[Survey]**

**Sanctions and restricted regions [Vendor policy, primary]**

- US-sanctioned regions named: Crimea, separatist Donetsk and Luhansk, Cuba, Iran, North Korea. GitHub is unavailable to developers in North Korea; Iran is served under an OFAC licence including paid services; Cuba's cloud services are generally available — [GitHub Trade Controls](https://docs.github.com/en/site-policy/other-site-policies/github-and-trade-controls)
- For restricted organizations: private repository access suspended, paid Actions minutes suspended, paid organization public repos may become read-only archives. In Crimea, Donetsk, and Luhansk, free users keep public repos and free Actions minutes for personal use. **Downloading or deleting private repository data is not allowed unless the US government authorizes it** — [GitHub Trade Controls](https://docs.github.com/en/site-policy/other-site-policies/github-and-trade-controls)
- The page describes no general GitHub.com restriction for Russia, Belarus, or Syria, but says GitHub Enterprise Server may not be sold or exported to Country Group E:1, listing Cuba, Iran, North Korea, Syria, Russia, Belarus — [GitHub Trade Controls](https://docs.github.com/en/site-policy/other-site-policies/github-and-trade-controls)

### Inferences
Decision pattern drawn from the above. This is my synthesis, not survey evidence:

| Situation | Likely default | Why |
|---|---|---|
| Code already on GitHub, no residency limits | GitHub Actions | SCM gravity, low ops load, highest recommend score |
| Mature Jenkins estate with scripted pipelines, many plugins | Keep Jenkins, migrate selectively | Importer leaves scripted pipelines and unknown plugins manual; 62% do not plan to switch |
| Special hardware, mainframe, build farm | Self-hosted runners on Actions, or Jenkins | Agent control matters; Bankdata shows both are viable |
| Air-gapped, strict jurisdiction, or GitHub export-restricted | Gitea, Forgejo, GitLab self-managed, or Jenkins | GHES unavailable in some regions; mirror actions internally |
| Small team on Gitea already | Gitea Actions | Familiar syntax; accept compat gaps and self-hosted runners |

- The trade-control exposure is the strongest sovereignty argument because the consequence is loss of access to private code, with an export block. It is a low-probability, high-impact risk for most, and an existing one for entities tied to those regions.
- EU residency addresses data location, not jurisdiction. That is why the Dutch and individual arguments, which focus on US law, are not answered by it.
- Sovereignty looks significant in European public sector and among a vocal engineering minority, but there is no evidence yet of measurable private-sector share loss for GitHub: GitHub Actions rose six points in JetBrains' 2026 data.

### Gaps
- **No survey quantifies sovereignty as a CI/CD decision criterion**, nor senior-engineer decision criteria generally. Any percentage claiming this should be treated as unverified.
- Whether GitHub-hosted runners execute in-region under EU residency: not confirmed in sources found. Check current GitHub documentation.
- Other European forges (Schleswig-Holstein, code.gouv.fr, code.europa.eu) appeared in a search summary without a source I could confirm, and the code.overheid.nl primary site was not read.
- The ICC and Microsoft email-account incident, often used in this debate, was not confirmed by my search; do not rely on it without separate verification.
- Practical accounts from Russia, Iran, or similar regions moving from GitHub to Gitea or Forgejo: none reliable found. Whether the trade-controls page is fully current for Syria is unknown.

## What alternatives are senior engineers choosing instead?

### Takeaway
GitLab CI is the only alternative with large measured share; Argo is significant but in a Kubernetes-centred sample and mostly serves a different role. Tekton is flat and small, and Buildkite and Dagger have no independent adoption data in what I found.

### Cited Findings
- **GitLab CI:** 32% of organizations in 2026, down from 34% — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/); 43% in the cloud-native CNCF 2024 sample — [CNCF PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en) **[Survey]**
- **Argo:** 45% (CNCF 2024) — [CNCF PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en). Argo CD is in the radar's "adopt" ring with 87% likely to recommend — [CNCF Radar](https://www.cncf.io/wp-content/uploads/2026/03/Q1-2026-CNCF-Technology-Radar-Report.pdf). A July 2025 CNCF end-user survey reports nearly 60% of respondents' Kubernetes clusters use Argo CD, with 97% of respondents running it in production (secondary reports) — [Efficiently Connected](https://www.efficientlyconnected.com/gitops-in-production-as-argo-cd-becomes-a-kubernetes-default/); [Computer Weekly](https://www.computerweekly.com/blog/Open-Source-Insider/CNCF-Argo-CD-is-top-of-the-GitOps-pops-for-Kubernetes). The CNCF share figure groups "Argo" and does not separate Workflows from CD.
- **Tekton:** 5% in both 2023 and 2024 — [CNCF PDF](https://www.linuxfoundation.org/hubfs/Research%20Reports/cncf_annual_survey24_031225a.pdf?hsLang=en); "trial" ring, 85% likely to recommend — [CNCF Radar](https://www.cncf.io/wp-content/uploads/2026/03/Q1-2026-CNCF-Technology-Radar-Report.pdf) **[Survey]**
- **Buildkite [Vendor]:** claims against GitHub Actions include a 256-job matrix cap per run, slow self-hosted startup, frequent reliability issues, no central governance view, and lacking queue priorities. Marketing material; unverified — [Buildkite](https://buildkite.com/docs/pipelines/advantages/buildkite-vs-gha). JetBrains' 2026 page names Buildkite but shows no 2026 figure — [JetBrains](https://www.jetbrains.com/pages/state-of-cicd2026/)
- **Third-party runners:** Depot's co-founder argues GitHub is extracting revenue from a slow, unreliable service, and he has a commercial interest — [The Stack](https://www.thestack.technology/github-to-charge-self-hosted-actions-users-what-does-that-look-like/) **[Vendor]**

### Inferences
- Many senior teams are probably combining rather than replacing: one CI tool plus Argo CD for delivery. Multi-tool rates and Argo's numbers are consistent with this, but no source I read cross-tabulates it.
- Tekton's flat 5% suggests it is a niche platform-building choice rather than a mainstream Jenkins replacement.

### Gaps
- **Dagger:** no independent adoption data or first-party migration account found.
- **Argo Workflows** on its own: no separate share figure found.
- **Buildkite:** no independent share data or verified first-party migration with metrics found.
