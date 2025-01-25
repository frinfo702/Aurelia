document.addEventListener('DOMContentLoaded', () => {
    const keywords = [
        "Tokyo", "Osaka", "remote", "React", "Golang",
        "Python", "Kubernetes", "AI", "cloud", "AWS",
        "intern", "full-time"
    ];
    const floatingKeywordsContainer = document.getElementById('floating-keywords');
    if (!floatingKeywordsContainer) return; // 安全策

    const numKeywords = 12; // Number of floating keywords
    const keywordElements = [];

    // Create keyword elements
    for (let i = 0; i < numKeywords; i++) {
        const keyword = keywords[Math.floor(Math.random() * keywords.length)];
        const elem = document.createElement('div');
        elem.classList.add('keyword');
        elem.textContent = keyword;
        floatingKeywordsContainer.appendChild(elem);
        keywordElements.push(elem);
    }

    // Initialize positions
    keywordElements.forEach(elem => {
        elem.style.top = `${Math.random() * 100}%`;
        elem.style.left = `${Math.random() * 100}%`;
        elem.style.opacity = 0.4;
    });

    let mouseX = window.innerWidth / 2;
    let mouseY = window.innerHeight / 2;

    document.addEventListener('mousemove', (e) => {
        mouseX = e.clientX;
        mouseY = e.clientY;
    });

    function animate() {
        keywordElements.forEach((elem, index) => {
            const speed = (index + 1) * 0.7;
            const targetX = mouseX + (Math.random() - 0.5) * 150;
            const targetY = mouseY + (Math.random() - 0.5) * 150;
            const currentX = parseFloat(elem.style.left);
            const currentY = parseFloat(elem.style.top);
            const newX = currentX + (targetX / window.innerWidth * 100 - currentX) * 0.05 * speed;
            const newY = currentY + (targetY / window.innerHeight * 100 - currentY) * 0.05 * speed;

            elem.style.transform = `translate(${newX}%, ${newY}%) scale(1.2)`;
            elem.style.opacity = 0.3 + Math.random() * 0.4;
        });
        requestAnimationFrame(animate);
    }

    animate();
});
