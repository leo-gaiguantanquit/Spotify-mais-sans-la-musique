// change le thème entre clair et sombre

(function () {
	document.addEventListener('DOMContentLoaded', function () {
		const checkbox = document.querySelector('.nav__links input[type="checkbox"]');
		if (!checkbox) return;

		// charge le thème sauvegardé ou celui par défaut du système
		const saved = localStorage.getItem('theme'); // 'dark' or 'light'
		let isDark;
		if (saved) {
			isDark = saved === 'dark';
			document.getElementById('nav__logo').src ="../static/img/DarkLogo.png";
		} else {
			isDark = !window.matchMedia || !window.matchMedia('(prefers-color-scheme: light)').matches;
			document.getElementById('nav__logo').src ="../static/img/LightLogo.png";

		}

		// applique le thème
		document.documentElement.classList.toggle('theme-light', !isDark);
		checkbox.checked = isDark;

		checkbox.addEventListener('change', function () {
			const nowDark = checkbox.checked;
			if (nowDark) {
				document.documentElement.classList.remove('theme-light');
				localStorage.setItem('theme', 'dark');
                document.getElementById('nav__logo').src ="../static/img/DarkLogo.png";
			} else {
				document.documentElement.classList.add('theme-light');
				localStorage.setItem('theme', 'light');
                document.getElementById('nav__logo').src ="../static/img/LightLogo.png";
			}
		});
	});
})();
