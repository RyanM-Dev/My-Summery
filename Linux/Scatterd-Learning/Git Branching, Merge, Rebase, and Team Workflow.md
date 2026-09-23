# 🌿 

A good Git workflow keeps features and bug fixes isolated, makes code review easier, and keeps `develop` and `main` stable.

A practical flow looks like this:

```text
feature/* ─────┐
bugfix/* ──────┤
               ▼
            develop
               │
               ▼
        release/v1.x.x
               │
               ▼
             main
```

The basic idea is:

```text
Feature / Bugfix
      ↓
Pull Request / Merge Request
      ↓
develop
      ↓
release branch
      ↓
main
      ↓
Version / Tag
```

---

# 🌿 One Branch per Feature or Bug Fix

Do not develop unrelated changes directly on `develop` or `main`.

Create a separate branch for each piece of work.

Examples:

```text
feature/user-login
feature/product-search
feature/payment-system

bugfix/user-auth-error
bugfix/payment-timeout
bugfix/product-price
```

For example:

```bash
git switch develop
git pull
git switch -c bugfix/user-auth-error
```

Now all commits related to the authentication bug stay isolated inside:

```text
bugfix/user-auth-error
```

This makes the branch easier to:

- review
    
- test
    
- merge
    
- revert
    
- understand later
    

> 💡 **Real-World Example**
> 
> Imagine two developers are working simultaneously:
> 
> ```text
> Developer A → feature/user
> Developer B → feature/product
> ```
> 
> Their unfinished work does not interfere with each other because each feature has its own branch.

---

# 🚀 Pushing a New Branch

When you create a branch locally:

```bash
git switch -c feature/user
```

the branch does not automatically exist on the remote repository.

If you run:

```bash
git push
```

Git usually tells you that the branch has no upstream branch and gives you a command similar to:

```bash
git push --set-upstream origin feature/user
```

or the shorter form:

```bash
git push -u origin feature/user
```

### What `-u` means

`-u` means:

```text
--set-upstream
```

It connects:

```text
local feature/user
        │
        ▼
origin/feature/user
```

After that, you can normally use:

```bash
git push
```

and:

```bash
git pull
```

without specifying the branch every time.

---

# 🔄 Our Development Flow

Features and bug fixes first go into `develop`.

For example:

```text
feature/user ──────┐
feature/product ───┤
bugfix/auth ───────┤
                   ▼
                develop
```

At the end of the sprint, a release branch is created from `develop`.

For example:

```text
release/v1.4.0
```

The flow becomes:

```text
feature/*
bugfix/*
    │
    ▼
 develop
    │
    ▼
release/v1.4.0
    │
    ▼
  main
```

A more complete history might look like:

```text
feature/user ───────┐
                    │
feature/product ────┼──▶ develop
                    │
bugfix/auth ────────┘
                         │
                         ▼
                  release/v1.4.0
                         │
                         ▼
                       main
                         │
                         ▼
                      v1.4.0
```

---

# 🏷️ Release Branches and Versioning

When the sprint is ready for release, create a versioned release branch.

For example:

```bash
git switch develop
git pull
git switch -c release/v1.4.0
```

The release branch can be used for final:

- testing
    
- release-specific fixes
    
- version changes
    
- documentation changes
    
- deployment preparation
    

When everything is ready:

```text
release/v1.4.0
      │
      ▼
     main
```

The final release can then be tagged:

```bash
git tag v1.4.0
git push origin v1.4.0
```

A tag gives an exact Git commit a meaningful version:

```text
main
  │
  ▼
commit abc123
  │
  └── v1.4.0
```

---

# 🔎 Pull Request / Merge Request Review

Before a feature or bug fix is merged into `develop`, it should go through a:

```text
Pull Request (PR)
```

or:

```text
Merge Request (MR)
```

depending on the Git hosting platform.

The workflow is:

```text
Developer implements feature
        ↓
Push branch
        ↓
Create PR / MR
        ↓
Reviewer checks code
        ↓
Reviewer leaves comments
        ↓
Developer fixes the code
        ↓
Push new commits
        ↓
Review again
        ↓
Merge
```

An important team rule is:

> When reviewing another developer's code, leave comments explaining the problem instead of silently changing their implementation yourself.

For example, instead of directly rewriting:

```cpp
if (user != nullptr) {
    login(user);
}
```

a reviewer might comment:

```text
Can this condition be moved into authenticateUser() so authentication
logic stays in one place?
```

The original developer then implements the change.

This helps the developer understand:

- what was wrong
    
- why it should change
    
- how the team expects the code to be structured
    

Code review therefore becomes part of the learning process, not only a gate before merging.

---

# ⚔️ Git Merge Conflicts

A Git merge conflict happens when Git cannot automatically decide which change should be kept.

Assume these branches exist:

```text
develop
feature/user
feature/product
```

Both feature branches originally came from the same `develop` commit:

```text
A---B   develop
     \
      feature/user
      feature/product
```

At commit `B`, suppose the code contains:

```go
timeout := 30
```

---

## 1️⃣ `feature/user` Changes the Code

Inside:

```text
feature/user
```

the developer changes:

```go
timeout := 30
```

to:

```go
timeout := 60
```

Then the branch is merged into `develop`.

```bash
git switch develop
git merge feature/user
```

There is no conflict because `develop` had not independently changed the same code.

Now:

```text
develop:

timeout := 60
```

Conceptually:

```text
A---B---------U---M   develop
     \
      P               feature/product
```

---

## 2️⃣ `feature/product` Also Changed the Same Code

Meanwhile, `feature/product` was created from the older commit `B`.

It changed:

```go
timeout := 30
```

to:

```go
timeout := 120
```

Now Git effectively sees:

```text
Common ancestor:

timeout := 30


develop:

timeout := 60


feature/product:

timeout := 120
```

Both branches changed the same original line differently.

---

## 3️⃣ `feature/product` Tries to Merge

Eventually:

```bash
git switch develop
git merge feature/product
```

Git compares three versions:

```text
Common ancestor
develop
feature/product
```

It sees:

```text
30 → 60

and

30 → 120
```

Git cannot know which one is logically correct.

Therefore it produces a conflict:

```go
<<<<<<< HEAD
timeout := 60
=======
timeout := 120
>>>>>>> feature/product
```

A developer must decide the final code.

For example:

```go
timeout := 120
```

Then mark the conflict as resolved:

```bash
git add .
git commit
```

---

# 🧠 Why Three Versions Matter

A conflict is easier to understand when you realize Git is not simply comparing two files.

Git usually considers:

```text
        Common Ancestor
             │
       ┌─────┴─────┐
       ▼           ▼
   Branch A     Branch B
```

Example:

```text
Original:

timeout = 30
```

Then:

```text
feature/user:

30 → 60
```

and:

```text
feature/product:

30 → 120
```

Because both branches changed the same original content differently, Git needs a human decision.

---

# ✅ Changes That Usually Do Not Conflict

Suppose one branch changes:

```go
timeout := 30
```

while another changes:

```go
maxRetries := 5
```

If these changes do not overlap, Git can usually combine them automatically:

```go
timeout := 60
maxRetries := 10
```

So a merge conflict does **not** simply mean:

```text
two developers changed the same file
```

Two developers can modify the same file without creating a conflict.

The important issue is usually:

```text
overlapping incompatible changes
```

---

# 🌐 Local and Remote Branches

Git usually has both local and remote-tracking branches.

For example:

```text
develop
```

is your local branch.

While:

```text
origin/develop
```

represents your local knowledge of the remote `develop` branch.

Conceptually:

```text
Your computer

develop
feature/user

       │
       │ fetch / pull / push
       ▼

Remote repository

origin/develop
origin/feature/user
```

---

# 🚫 Why `git push` Can Be Rejected

Imagine both your computer and the remote repository start here:

```text
A---B
```

You create a commit locally:

```text
A---B---L
        local
```

But another developer pushes first:

```text
A---B---R
        remote
```

Now the histories have diverged:

```text
        L   local
       /
A---B
       \
        R   remote
```

If you try:

```bash
git push
```

Git may reject the push because pushing `L` would not simply move the remote branch forward.

The remote contains commit `R`, which your local branch does not yet contain.

You first need to integrate the remote changes.

---

# 🔀 Option 1: `git pull` with Merge

You can run:

```bash
git pull
```

Conceptually, Git performs:

```text
git fetch
+
git merge
```

Before:

```text
        L
       /
A---B
       \
        R
```

After:

```text
        L
       / \
A---B     M
       \ /
        R
```

`M` is a merge commit.

Then you can push:

```bash
git push
```

---

# 🧹 Option 2: `git pull --rebase`

A cleaner alternative for many personal feature branches is:

```bash
git pull --rebase
```

The short form is:

```bash
git pull -r
```

Instead of creating a merge commit, Git fetches the remote commits and replays your local commits on top.

Before:

```text
        L
       /
A---B
       \
        R
```

After rebase:

```text
A---B---R---L'
```

Your original commit:

```text
L
```

is recreated as:

```text
L'
```

because it now has a different parent.

The history stays linear:

```text
A → B → R → L'
```

instead of:

```text
      L
     / \
A---B   M
     \ /
      R
```

---

# 💻 `git pull -r`

### What it does

```bash
git pull -r
```

is equivalent to:

```bash
git pull --rebase
```

Conceptually:

```text
fetch remote changes
        ↓
update your base
        ↓
temporarily remove your local commits
        ↓
apply remote commits
        ↓
replay your local commits on top
```

Example:

Before:

```text
Remote:

A---B---C


Local:

A---B---D---E
```

After:

```bash
git pull -r
```

you get:

```text
A---B---C---D'---E'
```

Instead of creating a merge commit.

---

# 🔀 Git Merge vs Rebase

Assume:

```text
A---B---C---D   develop
     \
      E---F     feature/product
```

`develop` has moved forward while the feature branch still comes from the older commit `B`.

---

## 🔀 Using Merge

From the feature branch:

```bash
git switch feature/product
git merge develop
```

Git combines the two histories.

```text
A---B---C---D
     \       \
      E---F---M
```

`M` is a merge commit.

Merge:

- preserves existing commit history
    
- does not rewrite `E` or `F`
    
- creates a merge commit
    
- clearly shows that two histories were joined
    

---

# 🧹 Using Rebase

Instead:

```bash
git switch feature/product
git rebase develop
```

Git temporarily takes:

```text
E
F
```

and replays them after `D`.

Before:

```text
A---B---C---D
     \
      E---F
```

After:

```text
A---B---C---D---E'---F'
```

Notice:

```text
E  → E'
F  → F'
```

These are new commits.

Their changes may be the same, but their commit hashes change because their history changed.

---

# 🧠 Rebase Does Not Replace `develop`

One important point:

```bash
git switch feature/product
git rebase develop
```

does **not** erase or replace `develop`.

`develop` remains:

```text
A---B---C---D
```

Git changes the feature branch:

```text
Before:

A---B
     \
      E---F


develop:

A---B---C---D
```

into:

```text
A---B---C---D---E'---F'
                \
                 feature/product
```

Think of rebase as:

> Move my work so that it looks like I started from the latest version of `develop`.

---

# ⚔️ Rebase Conflicts

Rebase applies your commits one by one.

Suppose your feature has:

```text
E
F
G
```

Git starts replaying them:

```text
E ✅
F ❌ conflict
G not applied yet
```

The rebase pauses.

You resolve the file manually.

Then:

```bash
git add .
```

and continue:

```bash
git rebase --continue
```

This tells Git:

> The current conflict has been resolved. Continue replaying the remaining commits.

Git continues:

```text
E' ✅
F' ✅
G' ✅
```

If another conflict appears during `G`, Git pauses again.

Resolve it:

```bash
git add .
git rebase --continue
```

Repeat until the rebase finishes.

---

# 🛑 Cancel a Rebase

If the rebase becomes confusing or you want to return to the original state:

```bash
git rebase --abort
```

Git restores the branch to how it looked before the rebase started.

This is extremely useful when learning rebase.

---

# ⚠️ Rebase and Shared Branches

Because rebase recreates commits, commit hashes change.

For example:

```text
E → E'
F → F'
```

Therefore rebasing a branch that multiple developers are already using can cause unnecessary confusion.

A useful rule is:

```text
Personal feature branch
        ↓
Rebase is usually fine
```

while:

```text
main
develop
shared branches
        ↓
Avoid rewriting published history
```

For your own feature branch, this workflow is useful:

```bash
git switch feature/product
git fetch origin
git rebase origin/develop
```

This updates your feature branch using the latest remote `develop`.

---

# 🔄 `git fetch` vs `git pull`

These commands are related but not identical.

|Command|Downloads Remote Changes|Changes Current Branch|
|---|--:|--:|
|`git fetch`|✅|❌|
|`git pull`|✅|✅|
|`git pull --rebase`|✅|✅|

### `git fetch`

```bash
git fetch origin
```

Updates your knowledge of branches such as:

```text
origin/develop
origin/main
```

but does not immediately modify your current branch.

This is useful because you can inspect remote changes first.

---

### `git pull`

Usually behaves roughly like:

```text
git fetch
+
git merge
```

---

### `git pull --rebase`

Behaves roughly like:

```text
git fetch
+
git rebase
```

---

# 🌍 Practical Feature Workflow

A normal task might begin with:

```bash
git switch develop
git pull
git switch -c feature/user-auth
```

Work normally:

```bash
git add .
git commit -m "Add user authentication"
```

Push the new branch:

```bash
git push -u origin feature/user-auth
```

Create a PR/MR:

```text
feature/user-auth
        ↓
     develop
```

If `develop` changes while you are working:

```bash
git fetch origin
git rebase origin/develop
```

If there is a conflict:

```bash
# fix files

git add .
git rebase --continue
```

Then push your branch.

If you previously pushed the branch before rebasing, the commit hashes have changed. In that situation, use Git's safer force-push form:

```bash
git push --force-with-lease
```

> ⚠️ Prefer `--force-with-lease` over plain `--force`. It adds protection against accidentally overwriting remote work you have not seen.

---

# 🧪 Interview Q&A

**Q1:** Why should each bug fix or feature have its own branch?

**Q2:** What does `git push -u origin feature/user` do?

**Q3:** Why can Git reject a push when another developer has pushed first?

**Q4:** What is the difference between `git pull` and `git pull -r`?

**Q5:** Why can two developers modify the same file without causing a conflict?

**Q6:** Why does rebase change commit hashes?

**Q7:** What does `git rebase --continue` do?

**Q8:** Why should published `main` or `develop` history generally not be rebased?

> [!answer]- 📋 Answers
> 
> **A1:** It isolates unrelated changes, making review, testing, merging, reverting, and debugging easier.
> 
> **A2:** It pushes the local branch to `origin` and configures the remote branch as its upstream, allowing future `git push` and `git pull` commands to work without repeatedly specifying the remote branch.
> 
> **A3:** The remote branch contains commits that the local branch does not have. Git prevents a normal push from overwriting that newer history.
> 
> **A4:** Normal `git pull` commonly integrates remote changes using a merge, while `git pull -r` rebases local commits on top of the updated remote branch and avoids an unnecessary merge commit.
> 
> **A5:** Git can often merge changes automatically when they affect different lines or non-overlapping sections. Conflicts usually occur when incompatible changes overlap.
> 
> **A6:** A commit hash depends partly on its parent commit. Rebase gives the commit a new parent, so Git creates a new commit with a different hash.
> 
> **A7:** It tells Git that the current rebase conflict has been resolved and that Git should continue replaying the remaining commits.
> 
> **A8:** Rebase rewrites commit history. Rewriting commits that other developers already use can make their local histories diverge and complicate collaboration.

---

# 🔨 Hands-On Practice

Create a test repository:

```bash
mkdir git-practice
cd git-practice

git init
```

Create an initial file:

```bash
echo "timeout=30" > config.txt

git add .
git commit -m "Initial config"
```

Create `develop`:

```bash
git switch -c develop
```

Create the first feature:

```bash
git switch -c feature/user
```

Change:

```text
timeout=30
```

to:

```text
timeout=60
```

Commit:

```bash
git add .
git commit -m "Change user timeout"
```

Merge it:

```bash
git switch develop
git merge feature/user
```

You can then create another branch from an older state and experiment with changing the same line differently to see how Git reports conflicts.

Useful commands while experimenting:

```bash
git status
git log --oneline --graph --all
git branch
```

A particularly useful visualization is:

```bash
git log --oneline --graph --decorate --all
```

It lets you actually see branches, merges, and rebases.

---

# 📋 Quick Reference

|Task|Command|
|---|---|
|Create feature branch|`git switch -c feature/name`|
|Create bugfix branch|`git switch -c bugfix/name`|
|Push new branch|`git push -u origin branch-name`|
|Download remote state|`git fetch origin`|
|Pull using merge|`git pull`|
|Pull using rebase|`git pull -r`|
|Rebase on latest develop|`git rebase origin/develop`|
|Mark conflict resolved|`git add .`|
|Continue rebase|`git rebase --continue`|
|Cancel rebase|`git rebase --abort`|
|Merge branch|`git merge branch-name`|
|View Git graph|`git log --oneline --graph --decorate --all`|
|Safer push after rebase|`git push --force-with-lease`|
|Create release branch|`git switch -c release/v1.4.0`|
|Create release tag|`git tag v1.4.0`|
|Push tag|`git push origin v1.4.0`|

---

# 🧠 Things to Remember

```text
One feature / bug
       ↓
One branch
```

```text
feature/*
bugfix/*
    ↓
 develop
    ↓
release/version
    ↓
   main
```

A merge conflict usually means:

```text
same original code
      ↓
changed differently
on different branches
      ↓
Git cannot safely choose
```

Merge:

```text
combines histories
```

Rebase:

```text
replays your commits on a newer base
```

And:

```bash
git pull -r
```

means:

```text
fetch remote changes
+
rebase local commits on top
```

---

# 💡 Pro Tips

Before starting a new feature, update `develop`:

```bash
git switch develop
git pull
```

Then create the branch:

```bash
git switch -c feature/my-feature
```

Before creating the final PR/MR, synchronize your branch with the latest remote `develop`:

```bash
git fetch origin
git rebase origin/develop
```

Use:

```bash
git status
```

whenever Git seems confusing. It usually tells you whether you are:

- merging
    
- rebasing
    
- resolving conflicts
    
- ahead of remote
    
- behind remote
    

And use:

```bash
git log --oneline --graph --decorate --all
```

whenever you want to understand what actually happened to the repository history.

During code review, explain problems through PR/MR comments and allow the original developer to implement the fix whenever practical. This keeps ownership clear and turns review into a learning process.

---

# 🔗 Related Topics

[[Git]]

[[Git Branches]]

[[Git Merge]]

[[Git Rebase]]

[[Git Merge Conflicts]]

[[Git Pull Requests]]

[[Git Tags]]

[[Semantic Versioning]]

[[CI-CD]]

git rm -r --chached . or filename : to remove file or dir from git after adding in git ignoe.
what does it do normaly?

git stash: to temprory hide changes to be able to switch to other branch without needing to commit or to hide recent changes (from the last commit) and check sth or check if the bug is caused by the recent changes.

git log: shows the history of commits, commits have commit hash that with git checkout hash we can checkout to that comit to check sth or even create a new branch out of that commit

git reset --hard HEAD~number of commits to be back to(1 or 5 oe...): is used when we want to reset changes to last or some privious commits. HEAD means last commit that can be shown in git log
