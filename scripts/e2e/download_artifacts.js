// .github/scripts/download-artifacts-and-get-pr.js
const core = require('@actions/core');
const github = require('@actions/github');
const fs = require('fs');
const path = require('path');
const AdmZip = require('adm-zip');

async function run() {
    try {
        const { owner, repo } = github.context.repo;
        const workflowRunId = github.context.payload.workflow_run.id;

        // List all artifacts from the triggering workflow run
        const artifacts = await github.getOctokit(process.env.GITHUB_TOKEN).rest.actions.listWorkflowRunArtifacts({
            owner,
            repo,
            run_id: workflowRunId,
        });

        // Download and extract each artifact
        for (const artifact of artifacts.data.artifacts) {
            const download = await github.getOctokit(process.env.GITHUB_TOKEN).rest.actions.downloadArtifact({
                owner,
                repo,
                artifact_id: artifact.id,
                archive_format: 'zip',
            });

            const zip = new AdmZip(Buffer.from(download.data));
            const extractPath = path.join(process.env.GITHUB_WORKSPACE, '.metrics', artifact.name);
            if (!fs.existsSync(extractPath)) {
                fs.mkdirSync(extractPath, { recursive: true });
            }
            zip.extractAllTo(extractPath, true);
            console.log(`Extracted artifact: ${artifact.name}`);
        }

        // Extract PR number
        let prNumber = null;
        const pullRequest = await github.getOctokit(process.env.GITHUB_TOKEN).rest.pulls.list({
            owner,
            repo,
            head: `${github.context.payload.workflow_run.head_repository.full_name}:${github.context.payload.workflow_run.head_branch}`,
        });

        if (pullRequest.data.length > 0) {
            prNumber = pullRequest.data[0].number;
        } else {
            // Fallback to commit SHA if needed
            const commitSha = github.context.payload.workflow_run.head_sha;
            const prsForCommit = await github.getOctokit(process.env.GITHUB_TOKEN).rest.repos.listPullRequestsAssociatedWithCommit({
                owner,
                repo,
                commit_sha: commitSha,
            });
            if (prsForCommit.data.length > 0) {
                prNumber = prsForCommit.data[0].number;
            }
        }

        if (prNumber) {
            console.log(`Found PR Number: ${prNumber}`);
            core.setOutput('pr_number', prNumber);
        } else {
            console.log('Could not determine PR number. Skipping comment.');
            core.setFailed('Could not determine PR number for commenting.');
        }
    } catch (error) {
        core.setFailed(error.message);
    }
}

run();