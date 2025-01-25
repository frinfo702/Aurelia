document.querySelector('.file-upload').addEventListener('click', () => {
    document.getElementById('resume').click();
});

document.getElementById('resume').addEventListener('change', function (e) {
    const fileName = e.target.files[0]?.name;
    if (fileName) {
        e.target.parentElement.querySelector('p').textContent = `Selected file: ${fileName}`;
    }
});

document.getElementById('applicationForm').addEventListener('submit', function (e) {
    e.preventDefault();
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;

    if (password !== confirmPassword) {
        alert('Passwords do not match');
        return;
    }

    // Here you would typically send the form data to your server
    alert('Application submitted successfully!');
});
