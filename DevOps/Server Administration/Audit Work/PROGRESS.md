# 📒 Audit Progress Log

See [[DevOps/Server Administration/Audit Work/PLAN|PLAN]] for the steps.

## 2026-11-11

- ✅ **Step 0 (before this plan existed):** Inspected `/opt/test-mqtt-5/deploy`, the second broker, ports and the NPM certificates for the MQTT broker guide. Several findings carry over: SSH passwords and root login enabled, ufw inactive, port 1884 plaintext, NPM admin on :81, the Gitea runner holds `docker.sock`, and `shop.example.net` serves the test UI over HTTP.

- ✅ **Step 1:** Surveyed the vault.
    - The vault root is `/home/ryan/Media/Code/Book/Obsidian`. Links are written from the root, like `DevOps/NaNa/...`.
    - Conventions: `AGENTS.md`, plus the Nana prompt as the style reference (concise, Mermaid, emojis, spacing, `> [!answer]-`).
    - The existing `DevOps/NaNa/Linux/7. Server Configuration.md` is a stub about non-root users, sudo and SSH keys. The new notes link back to it.
    - Books: *Linux Bible 11th (2026)*, *The Software Developer's Guide to Linux* (Packt), *CompTIA Linux+ Guide*, plus the `CompTIA Linux+` folder.
    - Decision: the new notes go in `DevOps/Server Administration/`. It is a sibling of `NaNa/` because this is Ryan's own server, not course material.

- 🔄 **Step 2: security audit**, batch 1 (accounts, SSH):
    - 22,941 failed SSH logins in 24 h. The server is under constant brute-force attack, with passwords allowed.
    - `/home/ryan/.ssh/authorized_keys` has mode **664** (group-writable).
    - `ubuntu` has `NOPASSWD` sudo (from cloud-init).
    - `/root/.ssh/authorized_keys` holds a `support@provider` key: the hosting provider can log in as root.
    - `x11forwarding yes`, `maxauthtries 6`, `logingracetime 120`.
    - SSH attack detail: 135 distinct IPs in 24 h; 8,780 password attempts against **root**; top usernames tried are `user`, `admin`, `debian`. Every *successful* login in the last 7 days was **root via public key** (36). `ryan` is never used. No fail2ban or CrowdSec. AppArmor is loaded (105 profiles enforcing).
    - Account status: `root` has a usable password (P); `ubuntu` is locked (L) but has NOPASSWD sudo; `ryan` has a password.

- 🔄 **Step 2: security audit**, batch 2 (Docker):
    - There is no `/etc/docker/daemon.json`. That means no global log rotation (most containers use json-file with **no size limit**), `live-restore=false`, and no userns-remap.
    - Hardening by container: the MQTT test stack, gitea and npm drop capabilities. `vending-machine-game` and the 3 `crypto-interop-spike` containers run as **root**, with no cap_drop, a writable rootfs and no limits. The spike containers also have `restart=no` (vending uses `unless-stopped`; corrected during review).
    - **No container except the test broker has a memory limit.** No swap is configured.
    - `gitea-runner` mounts `docker.sock`. Its config passes `--volume /opt/robot-market-front` into job containers, allows `valid_volumes` `/opt/test-mqtt-5/backend` and `/opt/robot-market-front`, and uses network `gitea_proxy`.
    - ⚠️ **`robot-market-mqtt5-test-broker-1` is UNHEALTHY.**
        - Every health check fails with `OCI runtime exec failed: procReady not received`. A manual `docker exec … true` fails the same way.
        - The cgroup shows `pids.current=61`, `pids.max=64` and `pids.events max 3,878,909`, but only **1** real task (mosquitto, 1 thread, 10 fds).
        - Hypothesis: leaked PID accounting from failed or killed execs, plus a very small `pids_limit: 64`. The cause is not confirmed.
        - The broker still listens and answers TLS on :8884. Its last log line is 2026-11-10 18:51 UTC, a scanner (`LZR`) probing the public plaintext port 1883/1884.

- ✅ **Steps 2–3: security, performance and reliability audit done.** Full evidence is in [[DevOps/Server Administration/Audit Work/RAW EVIDENCE|RAW EVIDENCE]]. New in batches 3–5:
    - Pending reboot: running kernel -30, installed -43, plus libc6. 72 days uptime.
    - World-readable `.env` files (644) everywhere under `/opt`.
    - Committed memory is 6.1 GB on 3.9 GB of RAM with no swap. IO PSI full is 6%.
    - The journal uses 3.6 GB with no cap. sshd produces most of the log lines.
    - 6 GB of Docker build cache can be reclaimed.
    - VS Code Server runs as root: 491 MB RAM and 4.5 GB on disk.
    - The only backup is one manual Gitea dump from 2026-09-10, on the same disk.
    - The NPM admin UI is reachable over :81 and via `npm.example.com`. The crypto-spike viewer on :8099 is public with no auth.
    - Most sites don't redirect HTTP to HTTPS, and no HSTS header was seen.
- 🔄 **Step 4:** mapping book chapters.
- ✅ **Step 4: book map** (page numbers are the books' **printed** page numbers, taken from each PDF's table of contents with `pdftotext`):
    - **Linux Bible 11th (2026):**
        - Ch 6 Processes p129
        - Ch 10 Software p217 (updates p233)
        - Ch 11 Users p247
        - Ch 12 Disks p271
        - Ch 13 Server Administration p305: secure the server p311, monitor p312, SSH p314, key-based auth p322, logging p324, sar p331, df/du p333
        - Ch 14 Networking p339
        - Ch 15 systemd services p367
        - Ch 21 Troubleshooting p505 (networking p520)
        - Ch 23 Basic Security p547: root access p549, passwords p551, dangerous permissions p558, updates p562, monitoring logs p563
        - Ch 24 Advanced Security p579: crypto p579, PAM limits p603
        - Ch 26 Securing Linux on a Network p641: nmap p642, firewalls p649, UFW p651
        - Ch 27 Clouds and Containers p663
    - **The Software Developer's Guide to Linux (Packt):**
        - Ch 2 Processes p21
        - Ch 3 systemd p37
        - Ch 7 Users and Groups p81
        - Ch 8 Ownership and Permissions p93
        - Ch 9 Installed Software p103
        - Ch 10 Configuring Software p115
        - Ch 13 SSH p167
        - Ch 15 Docker p203
        - Ch 16 Monitoring Application Logs p223
        - Ch 17 Load Balancing and HTTP p235
    - **CompTIA Linux+ and LPIC-1 Guide, 6th ed. (Eckert, 2024):**
        - Ch 9 Processes
        - Ch 10 Log File Administration p399 (journald p403)
        - Ch 11 System Backup p442
        - Ch 12 SSH p511
        - Ch 13 Containers p554
        - Ch 14 Security p583: securing the local computer p583, network attacks p591, performance monitoring p618, sysstat p619
    - ⚠️ The CompTIA 6th-edition PDF filename contains **newline characters**, so it cannot be linked reliably. The `( PDFDrive )` copy reports `0 page(s)` with `file`, so it may be damaged. The notes cite the CompTIA book by title and chapter only.
- 🔄 **Step 5:** writing the audit report.
- ✅ **Steps 5–7:** wrote the audit report (`1.`), concept notes `2.`–`11.`, and the learning path index (`0.`).
    - Linking from the Nana table of contents is still to do.
- 🔄 **Step 8:** critical review.
- ✅ Linked the learning path from the Nana table of contents, under the Linux table.
- ✅ **Step 8: critical review done.** See [[DevOps/Server Administration/Audit Work/REVIEW|REVIEW]].
    - 10 errors found and fixed. The most important: the claim that `authorized_keys` 664 blocks ryan's key was **wrong**; the key login works.
    - All book page numbers come from the tables of contents. 15 of them were spot-checked against the page bodies.
    - The link checker passes. The one broken link is **pre-existing** in the Nana TOC: `DevOps/NaNa/Linux/Bash Scripting/0. Bash Scripting Learning Path` does not exist.
- 🏁 **Done.** Ryan's next step: read the audit, then follow its Safe Fix Order, starting with SSH.
