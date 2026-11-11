# Jenkins as a CI/CD Platform in 2026: Senior DevOps Perspective

Research date: 2026-11-02. Source quality notes: jenkins.io and CloudBees docs were opened directly. Vendor pages (Harness, BuildPiper, Qovery, Northflank, HAMS Tech, cost-calculator sites) are labelled as such and are not neutral. Items marked "search snippet only" were not opened, so their details are unverified. The two Medium migration posts returned HTTP 403.

## What is the current architecture and the current LTS and Java baseline?

### Takeaway
As of 2 November 2026, the current LTS is 2.580.2 (28 October 2026), and it requires Java 21 or 25. Java 17 was dropped in 2.555.1 (April 2026). The basic design is still a stateful controller using `$JENKINS_HOME`, with builds offloaded to agents and Pipelines run through a Groovy CPS interpreter.

### Cited Findings

**Release state**
- Recent LTS releases: 2.580.2 (2026-10-28), 2.580.1 (2026-09-30), 2.568.3 (2026-09-02), 2.568.2 (2026-08-05), 2.568.1 (2026-07-08), 2.555.3 (2026-06-10), 2.555.2 (2026-05-13), 2.555.1 (2026-04-15), 2.541.3 (2026-03-18), 2.541.2 (2026-02-18), 2.541.1 (2026-01-21) — [LTS changelog](https://www.jenkins.io/changelog-stable/)
- 2.580.2 ships Jetty 12.1.14 and is tested on JDK 21 and 25 — [LTS changelog](https://www.jenkins.io/changelog-stable/)
- 2.568.1 lists Spring Security 7.0.5 to 7.1.0 and Spring Framework 7.0.7 to 7.0.8, while 2.555.1 lists Spring Framework 6.2.17 and Spring Security 6.5.9 — [LTS changelog](https://www.jenkins.io/changelog-stable/)
- 2.541.1 added a new GPG signing key for Linux repositories and raised the default maximum HTTP response header size to 32KB. 2.568.1 dropped Windows 2019 controller images. 2.555.2 changed the MSI signer to LF Open Source, LLC — [LTS changelog](https://www.jenkins.io/changelog-stable/)

**Java baseline**

| Java | First LTS | Final LTS | Planned end of support |
|---|---|---|---|
| 25 | 2.541.1 (Jan 2026) | current | Sept 2029 |
| 21 | 2.479.1 (Oct 2024) | current | Sept 2027 |
| 17 | 2.346.1 (June 2022) | 2.541.3 (Mar 2026) | unsupported |
| 11 | 2.164.1 (Mar 2019) | 2.462.3 (Oct 2024) | unsupported |

Source for the table: [Java support policy](https://www.jenkins.io/doc/book/platform-information/support-policy-java/)

- Installing an unsupported Java version means the controller will not run — [Java support policy](https://www.jenkins.io/doc/book/platform-information/support-policy-java/)

**Pipeline and Groovy CPS**
- CPS lets a Pipeline save its state to `program.dat` in the build directory and resume after a controller restart. Most Pipeline code runs in a special interpreter, which adds overhead — [CPS method mismatches](https://www.jenkins.io/doc/book/pipeline/cps-method-mismatches/)
- Compiled bytecode (JDK, core, plugins, Groovy runtime), constructors, and `@NonCPS` methods are not transformed. Non-CPS code cannot correctly call CPS code, and `@NonCPS` methods must not call steps such as `sh` or `node`. Passing a CPS closure to a library method such as a sort can silently return wrong results — [CPS method mismatches](https://www.jenkins.io/doc/book/pipeline/cps-method-mismatches/)

**Controller and agents**
- Jenkins docs call controller executors "generally a bad practice" because builds can overload CPU or memory and cause downtime. They recommend generic, interchangeable agents built from automation, images, or containers — [Hardware recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations/)
- Building on the controller gives build processes access to sensitive data in `$JENKINS_HOME` — [Architecting for scale](https://www.jenkins.io/doc/book/scaling/architecting-for-scale/)
- Each build node connection uses 2-3 threads and about 2 MB or more of memory — [Hardware recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations/)

**Modern reference pattern (vendor marketing)**
- HAMS Tech recommends JCasC for controller configuration, ephemeral Kubernetes agents, and Groovy Shared Libraries for pipeline logic. It claims controller recovery under 60 seconds with JCasC. The post (2026-07-28) has no named author and promotes the firm's managed DevOps pods, and I found no independent support for the 60-second figure — [HAMS Tech](https://www.hams.tech/blog/jenkins-modernization-myth-groovy-jcasc-2026.html)

### Inferences
- LTS releases in 2026 came exactly every 28 days, and a new baseline line arrived every 12 weeks (2.541, 2.555, 2.568, 2.580). Teams that take every security fix therefore need a roughly monthly upgrade slot.
- The changelog shows Spring 6 in 2.555.1 and Spring 7 in 2.568.1, so a Spring major-version move landed between those lines. That is a plausible plugin compatibility risk, but I found no report confirming breakage.
- Java 21 support ends in September 2027, so a Java 25 move is a planned 2027 task for sites on 21.
- The GPG key change is relevant to air-gapped mirrors, which must trust the new key.

### Gaps
- I did not open current JCasC, Kubernetes plugin, or Shared Libraries documentation. Their current versions and limitations are not covered here.
- The current weekly version is known only as at least 2.588, from the 2026-10-28 advisory.
- I did not review the Jenkins roadmap.

## What is the real maintenance burden?

### Takeaway
Monthly LTS patching is the baseline, and each extra controller adds its own plugin updates, backups, and security setup. Open-source Jenkins documentation points to splitting controllers for scale and resilience, while active/active HA is documented as a CloudBees feature.

### Cited Findings

**Scaling**
- Jenkins docs assume 5 CPU cores and one core per 100 jobs, or about 500 jobs per controller. Estimated jobs = developers x 3.333, controllers = jobs / 500, executors = jobs x 0.03. The worked example is 1,200 developers, about 4,000 jobs, about 8 controllers, and about 120 executors — [Architecting for scale](https://www.jenkins.io/doc/book/scaling/architecting-for-scale/) (undated page, references old versions such as 2.119 and 1.611, so treat as a rule of thumb)
- Vertical scaling is limited by hardware and makes each outage affect more teams. Horizontal scaling adds maintenance because each controller needs its own plugin updates, backups, and security setup. The docs suggest separate controllers for mission-critical projects — [Architecting for scale](https://www.jenkins.io/doc/book/scaling/architecting-for-scale/)
- Controller RAM ranges from about 200 MB for a small installation to 70+ GB for a very large controller — [Hardware recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations/)
- The docs recommend a larger disk over a faster one, saying Jenkins does not need fast disk access — [Architecting for scale](https://www.jenkins.io/doc/book/scaling/architecting-for-scale/)

**High availability**
- CloudBees CI active/active HA runs one logical controller across replicas, using a load balancer, a shared file system, and embedded Hazelcast. Other replicas take over a failed replica's builds — [CloudBees HA fundamentals](https://docs.cloudbees.com/docs/cloudbees-ci/latest/ha/ha-fundamentals)
- It has been available since CloudBees CI 2.414.2.2 (September 2023) — [CloudBees HA install guide](https://docs.cloudbees.com/docs/cloudbees-ci/2.414.2.2/ha-install-guide)
- Some Hazelcast upgrades prevent rolling upgrades, including 2.528.2.34846 and 2.568.1.37445. Certain Pipeline idioms and steps are incompatible with HA — [CloudBees HA fundamentals](https://docs.cloudbees.com/docs/cloudbees-ci/latest/ha/ha-fundamentals)
- Active/active on Windows controllers is unsupported in the 2.462.1.3 docs (2024-era version; current status not rechecked) — [CloudBees traditional HA install](https://docs.cloudbees.com/docs/cloudbees-ci/2.462.1.3/ha/specific-ha-installation-traditional)

**Staff time (mostly vendor estimates, no survey behind them)**
- See the cost section for the range, which runs from single-digit hours per month to 2-5 FTEs.

### Inferences
- Applying the docs' ratio, a 500-job estate needs about one controller, while a large organization may run many. Operating burden therefore grows with controller count, not just job count.
- From my background knowledge, not verified in this research: open-source Jenkins has no native active/active controller mode, so common options are active/passive failover, fast rebuild from configuration as code, and splitting controllers. Verify before stating this as fact.
- CPS durability adds a resume path after restarts, but the CPS restrictions make Shared Library code harder to write and debug than ordinary Groovy.

### Gaps
- I found no primary measurements of plugin upgrade breakage rates. Claims come from vendors such as Qovery and cicdcost.com.
- I did not find authoritative current JVM tuning guidance. The hardware page does not cover JVM settings, and I did not locate another page.
- I did not research backup procedures. The hardware page does not cover backups, and the CloudBees HA page does not cover backup and restore. Note that ThinBackup itself had a High-severity flaw (see security section).

## What does the security record look like in 2024-2026?

### Takeaway
Advisories are frequent and steady (11 in 2024, 12 in 2025, 10 so far in 2026), but the volume in autumn 2026 stands out: 68 CVEs across three advisories in eight weeks, including 6 Script Security sandbox bypasses and many core issues. The most serious historical event is CVE-2024-23897, which was exploited in ransomware attacks.

### Cited Findings

**Cadence**
- Jenkins published 33 advisories across 2024-2026: 11 in 2024, 12 in 2025, and 10 in 2026 through October 28 — [Advisory index](https://www.jenkins.io/security/advisories/)
- 2026 advisory dates: 02-18, 03-18, 04-29, 05-27, 06-10, 06-24, 08-05, 09-02, 09-16, 10-28 — [Advisory index](https://www.jenkins.io/security/advisories/)
- 6 of the 11 LTS releases in 2026 carried security fixes: 2.541.2, 2.541.3, 2.555.3, 2.568.2, 2.568.3, 2.580.2 — [LTS changelog](https://www.jenkins.io/changelog-stable/)

**CVE-2024-23897 (SECURITY-3314)**
- Rated Critical. The CLI's args4j `expandAtFiles` feature replaced `@path` with the file's contents, allowing arbitrary file reads on the controller. Users with Overall/Read could read whole files, and others could read the first few lines. Binary files could expose keys, enabling several routes to RCE. Affected weekly up to 2.441 and LTS up to 2.426.2; fixed in 2.442, 2.426.3, and 2.440.1 — [Advisory 2024-01-24](https://www.jenkins.io/security/advisory/2024-01-24/)
- CVE-2024-23898, published alongside it, was cross-site WebSocket hijacking on the CLI endpoint — [Advisory 2024-01-24](https://www.jenkins.io/security/advisory/2024-01-24/)
- CISA added CVE-2024-23897 to KEV on 2024-08-19, with a federal remediation deadline of 2024-09-09 and known ransomware use. CloudSEK and Juniper linked it to attacks on BORN Group and Brontoo Technology Solutions — [The Hacker News](https://thehackernews.com/2024/08/cisa-warns-of-critical-jenkins.html) (secondary source)

**Advisory 2026-09-02: 33 CVEs (CVE-2026-84645 to -84677)**
- 13 CVEs concern core or bundled Stapler, fixed in 2.580 and 2.568.3. These include CVE-2026-84645, an unsafe deserialization flaw that can lead to RCE via Script Console, plus "remember me" session fixation and CSRF token exposure — [Advisory 2026-09-02](https://www.jenkins.io/security/advisory/2026-09-02/)
- High-severity plugin issues include SAML (overwrite IdP metadata and authenticate as any user), Entra ID (privilege escalation), Performance (deserialization code execution), File Parameter (path traversal, possible RCE), and ThinBackup (arbitrary controller files included in backups) — [Advisory 2026-09-02](https://www.jenkins.io/security/advisory/2026-09-02/)
- Parameterized Remote Trigger stores tokens unencrypted in job `config.xml`, with no fix or workaround as of publication — [Advisory 2026-09-02](https://www.jenkins.io/security/advisory/2026-09-02/)
- A search summary suggested conflicting CVE mappings for SAML and ThinBackup. The official advisory resolves this: ThinBackup is CVE-2026-84667 and SAML is CVE-2026-84668.

**Advisory 2026-09-16: 20 CVEs (CVE-2026-92122 to -92141)**
- Script Security had 6 High-severity sandbox bypasses plus 2 other High issues, a classpath approval bypass and a TOCTOU race, all fixed in 1422.v06869826dd9b_ — [Advisory 2026-09-16](https://www.jenkins.io/security/advisory/2026-09-16/)
- 12 more issues affected plugins including Pipeline: Multibranch, Pipeline: Groovy Libraries, GitLab, Gradle, Coverage, and Keycloak Authentication — [Advisory 2026-09-16](https://www.jenkins.io/security/advisory/2026-09-16/)

**Advisory 2026-10-28: 15 CVEs, all core (13 High, 2 Medium)**
- Issues include XStream deserialization that can lead to RCE, a setup wizard race allowing unauthenticated admin creation, path traversal in the agent-loading endpoint with possible RCE, an artifact extraction race allowing arbitrary file writes, and "remember me" cookies persisting after API token revocation. Fixed in 2.588 and 2.580.2 — [Advisory 2026-10-28](https://www.jenkins.io/security/advisory/2026-10-28/)

**Earlier 2026 core**
- CVE-2026-33001 concerns symlink handling when extracting `.tar` and `.tar.gz` archives, fixed in 2.555 and 2.541.3 — [Advisory search result](https://cybersecuritynews.com/jenkins-vulnerabilities-expose-ci-cd-servers/) (search snippet only; attribution between the results is not confirmed)

### Inferences
- Script Security should not be treated as a hard security boundary: 6 bypasses in one advisory suggest that anyone who can edit Pipelines or Shared Libraries may be able to reach the controller if plugins lag. Agent isolation and restricting who can change pipeline code matter more.
- Several 2026 flaws let agent-supplied data attack the controller (agent-computed link XSS, artifact extraction race, agent-loading path traversal). Agents running untrusted PR code are therefore a real risk to the controller.
- The plugin model spreads risk across many maintainers. The unfixed Parameterized Remote Trigger case shows that an operator may have to remove a plugin rather than patch it.
- Advisories on identity plugins (SAML, Entra ID, Keycloak, LDAP) and on token-using plugins (GitLab SSRF using stored API tokens, XL Deploy credential ID enumeration) mean credentials and sign-in rely partly on plugin quality.
- The advisory cadence is similar year over year, but CVE volume per advisory looks high in late 2026. I did not count CVEs for earlier advisories, so this is an impression, not a measured trend.

### Gaps
- I found no authoritative explanation for the size of the September-October 2026 batches, such as a security audit or automated discovery.
- I have no CVE totals per year, no split of core versus plugin counts across the period, and no evidence on whether any 2025-2026 CVE was added to KEV.
- I did not read advisories 2026-08-05 or the others not listed above.
- I did not research the Credentials plugin's design, such as how secrets are encrypted at rest. The CVE-2024-23897 advisory does confirm that reading binary files could expose cryptographic keys.

## What are its unique strengths?

### Takeaway
The sources consistently point to control and self-hosting, with air-gapped and regulated environments as the main case. Evidence is mostly qualitative and often from vendors, and I found little first-hand reporting on hardware or embedded use.

### Cited Findings
- A JetBrains report puts Jenkins at 28% adoption, just behind GitHub Actions. This is reported on the Jenkins project's own blog by Kris Stern (2026-04-06), so it is not neutral, and I did not open the underlying report — [Jenkins blog](https://www.jenkins.io/blog/2026/04/06/jetbrains-report-highlights-jenkins-as-a-popular-tool-in-2026/)
- Northflank says teams behind firewalls or in isolated networks rely on Jenkins because cloud platforms are not an option, and that it remains the default in regulated, on-prem-only setups. Northflank is a competing vendor, and this is a search snippet only — [Northflank CircleCI comparison](https://northflank.com/blog/circleci-vs-jenkins); [Northflank GHA comparison](https://northflank.com/blog/github-actions-vs-jenkins)
- Northflank cites over 1,800 plugins, while the Jenkins blog says only "thousands". I could not confirm a count because stats.jenkins.io returned no data to the fetch tool — [Northflank](https://northflank.com/blog/github-actions-vs-jenkins); [Jenkins blog](https://www.jenkins.io/blog/2026/04/06/jetbrains-report-highlights-jenkins-as-a-popular-tool-in-2026/)
- Some teams keep Jenkins for enterprise release orchestration and use GitHub Actions for repository checks (individual Substack author, search snippet only) — [Arvind Verma](https://arvindverma021.substack.com/p/why-jenkins-still-rules-cicd-in-2026)
- SCM independence is visible in the plugin set: the September 2026 advisories cover plugins for GitLab, Bitbucket Server, Bitbucket Push and Pull Request, and Gitee — [Advisory 2026-09-02](https://www.jenkins.io/security/advisory/2026-09-02/); [Advisory 2026-09-16](https://www.jenkins.io/security/advisory/2026-09-16/)
- Agents can be built from pre-baked images, containers, or virtualization, and the controller design does not depend on a hosted runner fleet — [Hardware recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations/)

### Inferences
- Self-hosted Jenkins charges no per-minute fee, so cost scales with infrastructure and staff rather than usage. This matters most for long or heavy builds, and it is consistent with the cost-increase anecdote in the migration section.
- CPS durability and Groovy logic support orchestration that spans many steps and environments, but this is the same flexibility that creates the maintenance burden.
- Air-gapped operation carries its own work: mirroring plugins and update sites. The 2026-09-02 fix for `update-center2` XSS in self-hosted update sites shows that this path has its own patching needs.

### Gaps
- I found no first-hand 2025-2026 case studies on embedded, hardware-in-the-loop, or specific regulated-industry use. The strength is plausible but rests on vendor claims here.
- I did not verify the JetBrains report's methodology, or whether 28% measures any use or primary use.

## What is the true cost, and what is the commercial option?

### Takeaway
The software is free, and staff time is the main cost, but public estimates differ by a large factor and none cites a primary dataset. CloudBees CI is the main commercial route and adds active/active HA, but I found no public pricing.

### Cited Findings

**Vendor or calculator estimates (not neutral, no survey data)**
- Harness says Jenkins needs 2-5 FTEs for plugin maintenance. Harness sells a competing product (search snippet only) — [Harness](https://www.harness.io/comparison-guide/harness-ci-vs-jenkins)
- BuildPiper says 1-2 senior DevOps engineers full time for a 100-200 person team. BuildPiper sells a competing product (search snippet only) — [BuildPiper](https://buildpiper.io/campaign/hidden-cost-jenkins/)
- An eitt.academy guide estimates one DevOps engineer spends 20-30% of their time on Jenkins in a mid-size organization (search snippet only) — [eitt.academy](https://eitt.academy/knowledge-base/jenkins-vs-github-actions-vs-gitlab-ci-cicd-2026/)
- cicdcost.com headlines $80-300 per month in infrastructure plus 5-10 hours of maintenance (search snippet only). The same search also surfaced a $75 per hour loaded rate and $375-750 per month in staff time, but did not state which page these came from — [cicdcost.com](https://cicdcost.com/jenkins-self-hosted-cost)
- Local US or UK engineers cost $150,000-$200,000 per year, according to HAMS Tech, which sells managed pods at a claimed 50% of local cost (marketing) — [HAMS Tech](https://www.hams.tech/blog/jenkins-modernization-myth-groovy-jcasc-2026.html)

**CloudBees**
- CloudBees CI provides active/active HA, with rolling restarts, rolling upgrades, and auto-scaling described for modern cloud platforms. These are vendor docs, and limits are listed in the maintenance section — [CloudBees HA fundamentals](https://docs.cloudbees.com/docs/cloudbees-ci/latest/ha/ha-fundamentals)
- Search returned no CloudBees HA pricing, licensing tiers, or 2026 price changes — [search-based finding; CloudBees capability page](https://cloudbees.com/capabilities/continuous-integration/high-availability-horizontal-scalability)

### Inferences
- A defensible planning range is a meaningful fraction of one engineer for a single well-automated controller, rising to multiple engineers for many controllers, heavy plugin use, or strict patching requirements. This is my synthesis, not a measured number.
- The 2026 monthly LTS rhythm, plus off-cycle plugin advisories, makes the lowest estimates of a few hours per month look optimistic for a security-sensitive site. Each upgrade needs plugin updates and testing, and large advisory batches add urgency.

### Gaps
- I found no public CloudBees pricing. Treat it as quote-based until confirmed.
- I found no independent survey on Jenkins staffing or TCO.
- No infrastructure cost benchmarks at large scale were found.

## What do senior practitioners report in 2025-2026 about staying or leaving?

### Takeaway
The best-documented departure in this research is Slack (December 2024), driven by security issues, downtime, and poor UX. Migrations are real but costly, and they do not automatically save money. The case for staying is mostly about sunk cost, control, and in-place modernization, and much of that advocacy has a commercial interest.

### Cited Findings

**Leaving**
- Slack (Emmanuel Zerefa, 2024-12-16): after years of developer frustration with security issues, downtime, and poor UX, Slack moved a majority of its CI jobs to GitHub Actions. The scope for its tool was 242 pipelines across 3 Jenkins instances — [Slack Engineering](https://slack.engineering/migration-automation-easing-the-jenkins-%E2%86%92-gha-shift-with-help-from-ai/)
- GitHub Actions Importer fully converted about 50% of Slack's pipelines, failed on about 5%, and partly converted the rest. Eight unsupported steps and environments caused over 90% of full-conversion failures, and users had to rewire secrets stored outside GitHub — [Slack Engineering](https://slack.engineering/migration-automation-easing-the-jenkins-%E2%86%92-gha-shift-with-help-from-ai/)
- Slack projected saving over 1,300 hours, about 80% of the manual effort (1,700 hours down to about 360). These are projections, and developers later reported errors in generated workflows. The "100% accuracy" LLM result came from only about two dozen reviewed outputs — [Slack Engineering](https://slack.engineering/migration-automation-easing-the-jenkins-%E2%86%92-gha-shift-with-help-from-ai/)
- A Medium post (Surbhi, 2026) appears to report an eleven-month migration of 200+ pipelines saving about $40,000 per year, with at least two weeks of parallel running and an early decommission date. It returned 403 and the search summary did not clearly attribute these figures, so treat them as unverified — [Medium](https://medium.com/@surbhi19/jenkins-vs-github-actions-vs-gitlab-ci-in-2026-what-we-learned-migrating-200-pipelines-fb529fbd966b)
- A Harness case study reports that finding job owners was the main migration challenge. Harness is a vendor, and this is a search snippet only — [Harness](https://www.harness.io/blog/jenkins-to-cie-journey)

**Counterarguments and staying**
- A Medium post titled "We Migrated From Jenkins to GitHub Actions. CI Costs Went Up 5x." reportedly found slow builds came from poor pipelines rather than Jenkins, advised optimizing before migrating, and ended roughly where it started on cost. It returned 403, so this is unverified — [Medium](https://medium.com/engineering-playbook/we-migrated-from-jenkins-to-github-actions-ci-costs-went-up-5x-64a72eef68ce)
- HAMS Tech argues that replacing Jenkins causes a 12-18 month migration freeze and recommends in-place modernization. This is marketing for its managed DevOps pods — [HAMS Tech](https://www.hams.tech/blog/jenkins-modernization-myth-groovy-jcasc-2026.html)
- DevOps Oasis argues that years of build logic, credential patterns, and deployment workflows make switching costly (search snippet only) — [DevOps Oasis](https://devopsoasis.blog/jenkins-still-earns-its-rack-space-2/)
- CircleCI argued in 2021 that teams need not replace Jenkins where it works. This is dated vendor content, included only for context (search snippet only) — [CircleCI](https://circleci.com/blog/dont-migrate-from-jenkins/)
- A migration guide from JetBrains (2026-03) covers moving a single project to TeamCity, which suggests incremental migration is a sales path for vendors (search snippet only) — [JetBrains](https://blog.jetbrains.com/teamcity/2026/03/first-jenkins-teamcity-migration/)

**Claim to discount**
- An eitt.academy article says Jenkins on Kubernetes is the standard in 60% of new deployments. No source supports this, and it should not be repeated — [eitt.academy](https://eitt.academy/knowledge-base/jenkins-vs-github-actions-vs-gitlab-ci-cicd-2026/)

### Inferences
- Slack's numbers suggest a usable planning figure: roughly half of pipelines may convert automatically, with the hard part concentrated in a small set of unsupported steps and environments. A plugin and step inventory is therefore a sensible first step. This comes from one company in 2024, so it should not be generalized too far.
- Migration savings depend heavily on build minutes. Removing Jenkins maintenance can be offset by usage fees for heavy workloads, unless self-hosted runners are used.
- The recent stream of core and plugin CVEs supports the security argument for leaving, but it also shows active fixing. An unmaintained or lagging controller is the worst of the outcomes considered here.

### Gaps
- I found few named-company, first-hand reports from 2025-2026. Slack is from late 2024, and most newer material is vendor content or anonymous blogging.
- I found no conference talks with concrete data, and no named large-company account of deliberately staying on Jenkins with figures.
- I have no reliable data on the rate of migration away from Jenkins.
- The tool-call budget was exceeded before these gaps could be closed: 24 calls against a guideline of about 15. A follow-up pass could target Reddit and Hacker News threads, DevOpsDays or Jenkins-related talks, and named engineering blogs.
