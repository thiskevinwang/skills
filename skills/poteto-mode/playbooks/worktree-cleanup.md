### Worktree and simulator cleanup

1. Record disk usage. Run `scripts/worktree-audit.sh [-b <pr-base-ref>] [repo-path]`. It reads paths from Git and checks merge, branch, PR, and edit state. Its last-chat column is `unknown` unless `PSTACK_TRANSCRIPTS_DIR` names a transcript directory provided for this task. It never searches global Codex or Cursor chat storage.
2. Check the current Codex attachments and chat list. Find pinned or active chats that use each candidate or its child worktrees. A `verify-usage` bucket is a request to check usage, not permission to remove the worktree. Read only known chat IDs or task-provided transcripts.
3. Inspect tracked, untracked, and ignored files before removal. Confirm that needed work has a recoverable copy. Hold worktrees with active processes, unknown ownership, or unpreserved files. Ask only for a missing decision that the available evidence cannot settle.
4. For a managed Codex worktree, use `archive_worktree` with its exact attachment identity. The tool saves a recoverable snapshot. Preserve needed ignored files separately before archival. For an ordinary worktree, use `git worktree remove <path>` only within the authorized cleanup scope. Do not add `--force` or remove the directory by hand to bypass a refusal.
5. Re-read the worktree list and disk usage. Report the paths removed or archived, the measured space change, and why each held candidate needs to remain.
6. Remove simulator data or package caches only when the request covers them and the contents are confirmed disposable. Do not remove app chat storage as a cache cleanup step.

**Reply:** disk usage before and after, removed or archived worktrees, held paths and reasons.
