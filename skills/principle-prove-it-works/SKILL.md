---
name: principle-prove-it-works
description: "Apply after completing a task, before declaring done. Verify against the real artifact (run the feature, read the actual value, inspect the diff), not a proxy, self-report, or 'it compiles.'"
---

# Prove It Works

Verify every task output by checking the real thing directly. Do not infer from proxies, self-reports, or "it compiles."

**Why:** Unverified work has unknown correctness. Indirect verification (file mtimes, output freshness, agent self-reports, cached screenshots) feels cheaper than direct observation. Acting on a wrong inference costs far more than checking the source.

Check the real thing, not a proxy:
- Check process liveness directly, not indirectly through derived state
- Read the actual value, not a cached or derived representation
- When verification fails, suspect the observation method before suspecting the system

## Script the check when you can

The strongest proof is a deterministic script that re-runs the same comparison, not a one-time eyeball. Write the script, run it, and keep its output as an artifact a reviewer can re-run instead of trusting your word.

Keep the artifact visible for the human. Commit it only for large or complex work where the trail has to be auditable later, like a big port or migration (the **show-me-your-work** skill).

## Make the proof readable

When a human reviewer needs to assess a run, lead with a readable report. Keep JSON as supporting evidence. Show the exercised path, expected and observed outcomes, concrete artifacts, and limits.

For a storage claim, include an actual bounded read of the stored row or object, with its source and capture time. A success response or a row ID is not the stored result. Label redactions and distinguish captured evidence from a live run.

Use `$verification-report` when installed for multi-step runs or captured records that need a readable presentation. Otherwise write the report directly. Verify any rendered report against its source data before delivery.
