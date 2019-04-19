
{
	setTimeout(() => document.body.classList.add('render'), 60);
	const navdemos = document.querySelectorAll('nav.demos > .demo');
	Array.from(navdemos).forEach(link => link.addEventListener('click', (ev) => {
		ev.preventDefault();
		console.log(ev.target.href);
		document.body.classList.remove('render');
		document.body.addEventListener('transitionend', () => window.location = ev.target.href);
	}));
}