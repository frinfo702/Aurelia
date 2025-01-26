// 求人が表示されない
document.addEventListener('DOMContentLoaded', async () => {
    const detailEl = document.getElementById('job-detail');
    const pathParts = window.location.pathname.split('/');
    const lastPart = pathParts[pathParts.length - 1];
    const parsedID = parseInt(lastPart, 10);

    let fetchURL = '/api/jobs';
    let isList = true;
    if (!isNaN(parsedID)) {
        fetchURL += `/${parsedID}`;
        isList = false;
    }

    try {
        const res = await fetch(fetchURL);
        if (!res.ok) {
            throw new Error('Failed to fetch job data');
        }
        const data = await res.json();

        if (isList) {
            // 複数求人表示
            detailEl.innerHTML = data.map(job => `
              <div>
                <h2>${job.technology_type} (${job.hiring_type})</h2>
                <p><strong>Job Tag:</strong> ${job.job_tag}</p>
                <p><strong>Income Range:</strong> ${job.income_range} yen</p>
                <p><strong>Requirements:</strong> ${job.requirements || 'Not specified'}</p>
                <p><strong>Used Tech:</strong> ${job.used_technology}</p>
                <hr/>
              </div>
            `).join('');
        } else {
            // 単一求人表示
            const job = data;
            detailEl.innerHTML = `
              <h2>${job.technology_type} (${job.hiring_type})</h2>
              <p><strong>Job Tag:</strong> ${job.job_tag}</p>
              <p><strong>Income Range:</strong> ${job.income_range} yen</p>
              <p><strong>Requirements:</strong> ${job.requirements || 'Not specified'}</p>
              <p><strong>Used Tech:</strong> ${job.used_technology}</p>
            `;
            // ヒーロー部分にタイトルを差し込む例
            const heroTitle = document.querySelector('.hero-job-title');
            const heroSub = document.querySelector('.hero-job-sub');
            if(heroTitle) heroTitle.textContent = `${job.technology_type} - ${job.hiring_type}`;
            if(heroSub) heroSub.textContent = job.job_tag || 'Top Skilled Position';
        }
    } catch (err) {
        console.error(err);
        detailEl.innerHTML = `<p style="color:red;">Error: ${err.message}</p>`;
    }
});
