// change le thème entre clair et sombre

(function () {
	function initTheme() {
		console.log("Theme script initializing...");
		const switches = document.querySelectorAll('.theme-switch');
		if (switches.length === 0) {
			console.warn("No theme switches found");
			return;
		}

		// charge le thème sauvegardé ou celui par défaut du système
		const saved = localStorage.getItem('theme'); // 'dark' or 'light'
		let isDark;
		if (saved) {
			isDark = saved === 'dark';
		} else {
			isDark = !window.matchMedia || !window.matchMedia('(prefers-color-scheme: light)').matches;
		}

		console.log("Initial theme isDark:", isDark);
		// applique le thème initial
		updateTheme(isDark);

		// ajoute les écouteurs sur tous les switchs
		switches.forEach(sw => {
			sw.addEventListener('change', function () {
				console.log("Switch changed to:", this.checked);
				updateTheme(this.checked);
			});
		});

		function updateTheme(dark) {
			console.log("Updating theme to dark:", dark);
			if (dark) {
				document.documentElement.classList.remove('theme-light');
				localStorage.setItem('theme', 'dark');
				const logo = document.getElementById('nav__logo');
				if(logo) logo.src = "/static/img/DarkLogo.png";
			} else {
				document.documentElement.classList.add('theme-light');
				localStorage.setItem('theme', 'light');
				const logo = document.getElementById('nav__logo');
				if(logo) logo.src = "/static/img/LightLogo.png";
			}
			
			// synchronise tous les boutons
			switches.forEach(sw => {
				sw.checked = dark;
			});
		}
	}

	if (document.readyState === 'loading') {
		document.addEventListener('DOMContentLoaded', initTheme);
	} else {
		initTheme();
	}
})();
