// ファイル例: frontend_mock/static/assets/js/jobs.js

document.addEventListener('DOMContentLoaded', async () => {
  const listEl = document.getElementById('jobs-list');
  try {
    const res = await fetch('/api/jobs');
    if (!res.ok) {
      throw new Error('Failed to fetch jobs');
    }
    const jobs = await res.json();

    if (!jobs || jobs.length === 0) {
      listEl.innerHTML = '<p>No jobs found.</p>';
      return;
    }

    let html = '';
    jobs.forEach(job => {
      // サーバーは "/jobs/{id:[0-9]+}" で詳細ページを返す想定
      html += `
        <div style="border:1px solid #ccc; padding:8px; margin:8px 0;">
          <h3>${job.technology_type} (${job.hiring_type})</h3>
          <p><strong>Used Tech:</strong> ${job.used_technology}</p>
          <p><strong>Income:</strong> ${job.income_range} yen</p>
          <!-- ここをクリックすると /jobs/1 などへ遷移し、detailページを返す -->
          <a href="/jobs/${job.job_id}" style="color:blue;text-decoration:underline;">View Detail</a>
        </div>
      `;
    });
    listEl.innerHTML = html;
  } catch (err) {
    console.error(err);
    listEl.innerHTML = `<p style="color:red;">Error: ${err.message}</p>`;
  }
});
