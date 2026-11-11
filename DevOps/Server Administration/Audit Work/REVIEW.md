# 🔍 Critical Review of This Audit

A second, skeptical pass over my own work: what I re-checked, what was **wrong** and got fixed, and what is still uncertain or missing.

See [[DevOps/Server Administration/Audit Work/PLAN|PLAN]], [[DevOps/Server Administration/Audit Work/PROGRESS|PROGRESS]] and [[DevOps/Server Administration/Audit Work/RAW EVIDENCE|RAW EVIDENCE]].

---

## ❌ Errors found and fixed

| # | What I wrote | What was true | How I found it | Fixed in |
|---|---|---|---|---|
| 1 | "`authorized_keys` 664 makes sshd **refuse** ryan's key. That probably explains why ryan never logs in." | A key login as `ryan` **works**. OpenSSH tolerates group-write when the group is the user's private group (`ryan:ryan`). | Actually tried it: `ssh -o BatchMode=yes ryan@… true` → `Accepted publickey for ryan` | Audit M4, SSH note, RAW EVIDENCE |
| 2 | Progress log: `vending-machine-game` has `restart=no` | It is `unless-stopped`. Only the 3 spike containers are `no`. | Re-read the container table | PROGRESS |
| 3 | Audit diagram: "shows **seven** doors" | Six arrows from the internet | Counted the edges | Audit |
| 4 | RAW EVIDENCE: `nft … grep -c rule = 0` presented as meaningful | That grep can't count nft rules (nft output never contains the word "rule"). The real evidence is `iptables -S DOCKER-USER` → empty. | Re-read the command | RAW EVIDENCE |
| 5 | Backups note: `gitea dump` run with `-u git` and no config path | The rootless image already runs as `git`. It needs `-c /etc/gitea/app.ini` and a writable working directory. | `docker inspect` / `id` in the container | Backups note |
| 6 | Backups note: `mariadb-dump -p<password>` | This contradicts the Secrets note ("never put passwords on the command line"). Now uses `MYSQL_PWD` from the container's own env (`MYSQL_ROOT_PASSWORD` verified by name only). | Consistency check across notes | Backups note |
| 7 | Secrets note: `passwd -S root ubuntu ryan` | `passwd -S` accepts **one** user. My own run of it printed the usage text. | Command output earlier in the audit | Secrets note |
| 8 | Docker note: "restart docker to apply daemon.json" with no ordering | `live-restore` can be applied with a **reload**, without stopping containers. Restarting first stops everything. | Re-checked the order of operations | Docker note |
| 9 | Firewall note: DOCKER-USER example without `--ctdir ORIGINAL` | Docker's documented pattern includes it, so reply packets are not matched | Compared against the Docker docs pattern | Firewall note |
| 10 | Memory note: IO pressure "likely" caused by journal writes | Not investigated. Reworded to "possible", with the tools to find the real cause. | Speculation check | Memory note |

**Lesson for me:** item 1 was a confident-sounding inference that I had not tested, and the test took 2 seconds. Every other "probably" in the notes was checked the same way, or marked as unverified.

---

## ✅ Verified on re-check

- 📚 **Book page numbers.** Every page number came from the books' own tables of contents.
    - Spot-checked against the page bodies: *Linux Bible* p. 324, 331, 549, 558, 649; *SDGL* p. 167; **all 9** CompTIA references (the PDF index is offset +16 from printed page numbers, consistently).
    - Not individually checked: the remaining *Linux Bible* and SDGL entries. Their headings wrap across lines in the PDF text, so they are taken from the TOC only.

- 🔗 **Links.** Every wikilink target exists. The first checker run showed "broken" links, but that was a bug in my checker: it didn't unescape `\|` in tables. In-note `#` anchors match real headings, and there are no unbalanced code fences.

- 🧾 **Facts behind the high-severity findings,** re-checked against command output:
    - the sshd drop-in order
    - the failed-login counts
    - the runner's `docker.sock` mount, and that `container.docker_host` is unset
    - `:81` and `:8099` returning 200 **from outside** (which proves there is no provider firewall in front of them)
    - the reboot-required package list
    - the backup directory dates
    - the `shop.example.net` page title (`MQTT test viewer`)

---

## ⚠️ Still uncertain

- 🧮 **The old broker's PID leak.** The facts are solid: `pids.current=61`, 1 real task, `pids.events max 3.88 M`, every exec failing. The *mechanism* (why PIDs are counted with no visible tasks) is a hypothesis. A restart will fix the symptom, but nobody knows whether it comes back. This is worth watching on the new broker, which uses `pids_limit: 64` too. Consider 256 there.

- 🏃 **act_runner mounting the socket into job containers.** This is based on act_runner's documented default for an unset `docker_host`. I did not run a test workflow to prove it. The runner itself definitely has the socket, so H1 stands either way.

- 🔐 **The `support@provider` key.** It is the provider's, by its comment. Whether you agreed to it is your call, not a finding of wrongdoing.

- 📊 **"Performance."** This is one snapshot (around 12:00 UTC on an idle weekday). PSI and `docker stats` say nothing about behaviour during CI builds or under load. A week of `sar` / sysstat data would be needed for real performance conclusions.

---

## 🕳️ Gaps: what this audit did NOT check

Be aware these exist. They are good follow-up exercises:

- 🌐 **IPv6 exposure.** Several services listen on `[::]`. I didn't check whether the VM has a public IPv6 address.
- 🧬 **CVEs in running images and app versions** (Gitea 1.27.3, NPM 2.15.1, MariaDB, Postgres 16). Next step: `trivy image`.
- 🧾 **NPM itself:** admin account strength and 2FA, and its access lists. I only looked at its generated nginx files.
- 🕵️ **Signs of past compromise:** unknown SUID files, unexpected cron jobs per user, rootkits, odd outbound connections. With password SSH open for months this is not paranoid. A basic check: `find / -xdev -perm -4000 -newer /etc/hostname`, user crontabs, `ss -tnp state established`.
- 🗄️ **Database settings:** Postgres and MariaDB configuration, and their users and grants.
- 📦 **Gitea:** its own backups, LFS and package storage growth, and the runner's registered repositories.
- 🖥️ **Kernel and sysctl hardening** beyond the basics (e.g. `kernel.unprivileged_userns_clone`, `net.ipv4.conf.*.accept_redirects`).

---

## 🎯 Is the output what Ryan asked for?

| Asked for | Delivered | Honest assessment |
|---|---|---|
| Analyze the server | 5 batches of read-only commands, with the evidence file | ✅ Broad. Gaps listed above. |
| Security and performance mistakes | 17 findings, ranked, each with evidence, fix and check | ✅ Security is strong. Performance is a single snapshot (see above). |
| Topics to learn, at his level | A 10-note path that builds on the Nana notes, plus "after this path" | ✅ Matches the course level; it doesn't jump to Kubernetes. |
| Obsidian references | Notes in the vault style (emojis, Mermaid, `[!answer]-` Q&A), linked from the Nana table of contents | 🟡 The notes are **longer** than the Nana prompt prefers ("shorter prose, compact exercises"). The concept notes could be trimmed after Ryan reads them. |
| Use and reference the books | Chapter and page references in every note, plus a book table | ✅ Verified as above. The CompTIA file can't be linked (newline in its filename). |
| Plan, progress, critical self-check | This folder | ✅ |

---

## 🧯 Side effects on the server

Being transparent about everything this work touched:

- 🐳 Pulled `eclipse-mosquitto:2.1.2-alpine` (new digest). This happened during the earlier MQTT guide work.
- 🧪 Ran short-lived, `--network none`, `--rm` test containers. Nothing persists.
- 📄 Copied small probe scripts to `/tmp` on the server and **deleted** them afterwards.
- 🔑 **One SSH login as `ryan`** (2026-11-11 12:19 UTC) to test the key. It is in the auth log as `Accepted publickey for ryan`.
- ⚙️ Everything else was read-only. No configuration, service, container or package was changed.

---

## 📝 Noticed in passing

- The Nana table of contents links to `DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path`, but that note doesn't exist (this predates this work). Either create it, or fix the link in the TOC and in `0. Linux Learning Path`.
