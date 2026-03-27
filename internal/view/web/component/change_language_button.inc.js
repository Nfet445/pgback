window.alpineChangeLanguageButton = function () {
  return {
    currentLang: window.currentLanguage || "en",

    setLanguage(lang) {
      this.currentLang = lang;
      window.setLanguage(lang);
      window.location.reload();
    },

    init() {
      this.currentLang = window.currentLanguage || "en";
    },
  };
};
