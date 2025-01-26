// language.js

// translations オブジェクトはそのまま
const translations = {
    ja: {
        // ...
    },
    en: {
        // ...
    }
};

function applyTranslations(lang) {
    const t = translations[lang] || translations.ja;
    // IDごとに書き換え
    document.getElementById('nav-jobs').textContent = t.nav_jobs;
    // ... 以下同様
}

function getLangFromQuery() {
    const params = new URLSearchParams(window.location.search);
    return params.get('lang') || 'ja';
}

function switchLang(lang) {
    const url = new URL(window.location);
    url.searchParams.set('lang', lang);
    window.location = url.toString();
}

document.addEventListener('DOMContentLoaded', () => {
    const currentLang = getLangFromQuery();
    applyTranslations(currentLang);
});
