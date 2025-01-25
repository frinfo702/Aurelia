// /static/assets/js/landing_page.js
// /api/jobs から取得して、先頭8件を .jobs-list に差し込む

document.addEventListener('DOMContentLoaded', async () => {
    const jobsListEl = document.getElementById('jobs-list');
    if (!jobsListEl) return;

    try {
        const res = await fetch('/api/jobs');
        if (!res.ok) {
            throw new Error('Failed to fetch jobs');
        }
        let jobs = await res.json();
        if (!Array.isArray(jobs)) jobs = [];

        // 8件だけ
        const topJobs = jobs.slice(0, 8);

        if (topJobs.length === 0) {
            jobsListEl.innerHTML = '<p>No jobs found.</p>';
            return;
        }

        let html = '';
        topJobs.forEach(job => {
            html += `
          <div class="job-card">
            <h3>${job.technology_type} (${job.hiring_type})</h3>
            <div class="job-detail">Used Tech: ${job.used_technology}</div>
            <div class="job-income">Income: ${job.income_range} yen</div>
            <a class="detail-link" href="/jobs/${job.job_id}">View Detail</a>
          </div>
        `;
        });

        jobsListEl.innerHTML = html;
    } catch (err) {
        console.error(err);
        jobsListEl.innerHTML = `<p style="color:red;">Error: ${err.message}</p>`;
    }
});
