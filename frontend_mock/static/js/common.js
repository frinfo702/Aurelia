// js/common.js

// 言語切り替え（簡易実装）
function switchLang(lang) {
    alert(`切り替え先言語: ${lang} (実際には再読み込み or 動的置換など実装)`);
    // 本格的にはページリロード or i18n対応
}

// ダークモード切り替え（簡易例）
function toggleDarkMode() {
    document.body.classList.toggle('dark-mode');
}
