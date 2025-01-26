// ジョブカードをクリックした時の処理を追加
document.addEventListener('DOMContentLoaded', () => {
    const jobCards = document.querySelectorAll('.job-card');
    jobCards.forEach(card => {
        card.addEventListener('click', () => {
            window.location.href = '/job-detail.html';
        });
    });
});
