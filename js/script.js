const labels = document.querySelectorAll('.circle-container label');

labels.forEach(label => {
	label.addEventListener('dblclick', () => {
		if (label.classList.contains('x')) {
			label.classList.remove('x');
			label.classList.add('o');
		} else {
			label.classList.remove('o');
			label.classList.add('x');
		}
	});
});

const form = document.querySelector('.user-profile');
const loginQuestion = document.querySelector('.login-question');
const registerQuestion = document.querySelector('.register-question');

registerQuestion.addEventListener('click', () => {
	form.classList.add('active');
});

loginQuestion.addEventListener('click', () => {
	form.classList.remove('active');
});

openBtn.addEventListener('click', () => {
	modal.show();
});

document.addEventListener('click', (e) => {
	const isClickOnOpenBtn = e.target === openBtn;
	if (isClickOnOpenBtn) return;
	const isClickOnDialog = e.target === modal;
	const isClickOnDialogContent = modal.contains(e.target);
	const isClickOutsideOfDialog = !(isClickOnDialog || isClickOnDialogContent);
	if (isClickOutsideOfDialog) {
		modal.close();
	}
});

let activeButton = document.getElementById('btn-one');
let saveButton = null;

function changeMode(button) {
	const activePage = document.querySelector('.content__page:not(.visually-hidden)');
	if (activePage) {
		activePage.classList.add('visually-hidden');
		// setCanvasSize();
	}

	if (activeButton) {
		activeButton.classList.remove('content__button--active');
		activeButton.classList.add('content__button--inactive');
	}

	activeButton = button;
	activeButton.classList.add('content__button--active');
	activeButton.classList.remove('content__button--inactive');

	const pageToShow = document.getElementById(button.dataset.target);
	if (pageToShow) {
		pageToShow.classList.remove('visually-hidden');
	}
}

function saveMode() {
	if (activeButton) {
		if (activeButton.classList.contains('content__button--save')) {
			activeButton.classList.remove('content__button--save');
			activeButton.classList.add('content__button--active');
		} else {
			activeButton.classList.remove('content__button--active');
			activeButton.classList.add('content__button--save');
		}
	}
}

// function setCanvasSize() {
// 	const canvas = document.querySelector('canvas');
// 	if (canvas) {
// 		canvas.width = document.body.clientWidth;
// 		canvas.height = document.body.clientHeight;
// 	}
// }

// (function () {
// 	var canvas = document.createElement('canvas'),
// 		ctx = canvas.getContext('2d'),
// 		particles = [],
// 		properties = {
// 			bgColor: 'rgba(0, 0, 0, 0)',
// 			particleImages: [],
// 			particleRadius: 14,
// 			particleCount: 40,
// 			particleMaxVelocity: 0.1,
// 		};

// 	function setCanvasSize() {
// 		canvas.width = document.body.clientWidth;
// 		canvas.height = document.body.clientHeight;
// 	}

// 	for (let i = 1; i <= 19; i++) {
// 		properties.particleImages.push(`./images/${i}.webp`);
// 	}

// 	properties.particleImages = properties.particleImages.map(src => {
// 		const img = new Image();
// 		img.src = src;
// 		return img;
// 	});

// 	canvas.style.position = 'fixed';
// 	canvas.style.top = '0';
// 	canvas.style.left = '0';
// 	canvas.style.zIndex = '-1';

// 	document.querySelector('body').appendChild(canvas);

// 	setCanvasSize();

// 	window.onresize = function () {
// 		setCanvasSize();
// 	}

// 	class Particle {
// 		constructor() {
// 			this.x = Math.random() * canvas.width;
// 			this.y = Math.random() * canvas.height;
// 			this.velocityX = Math.random() * (properties.particleMaxVelocity * 2) - properties.particleMaxVelocity;
// 			this.velocityY = Math.random() * (properties.particleMaxVelocity * 2) - properties.particleMaxVelocity;
// 			this.image = properties.particleImages[Math.floor(Math.random() * properties.particleImages.length)];
// 		}
// 		position() {
// 			this.x + this.velocityX > canvas.width && this.velocityX > 0 || this.x + this.velocityX < 0 && this.velocityX < 0 ? this.velocityX *= -1 : this.velocityX;
// 			this.y + this.velocityY > canvas.height && this.velocityY > 0 || this.y + this.velocityY < 0 && this.velocityY < 0 ? this.velocityY *= -1 : this.velocityY;
// 			this.x += this.velocityX;
// 			this.y += this.velocityY;
// 		}
// 		reDraw() {
// 			ctx.drawImage(this.image, this.x - properties.particleRadius, this.y - properties.particleRadius, properties.particleRadius * 2, properties.particleRadius * 2);
// 		}
// 	}

// 	function reDrawBackground() {
// 		ctx.fillStyle = properties.bgColor;
// 		ctx.fillRect(0, 0, canvas.width, canvas.height);
// 	}

// 	function reDrawParticles() {
// 		for (var i in particles) {
// 			particles[i].position();
// 			particles[i].reDraw();
// 		}
// 	}

// 	function loop() {
// 		ctx.clearRect(0, 0, canvas.width, canvas.height);
// 		reDrawBackground();
// 		reDrawParticles();
// 		requestAnimationFrame(loop);
// 	}

// 	function init() {
// 		for (var i = 0; i < properties.particleCount; i++) {
// 			particles.push(new Particle);
// 		}
// 		loop();
// 	}

// 	let imagesLoaded = 0;
// 	properties.particleImages.forEach(img => {
// 		img.onload = () => {
// 			imagesLoaded++;
// 			if (imagesLoaded === properties.particleImages.length) {
// 				init();
// 			}
// 		};
// 		img.onerror = () => {
// 			console.error(`Ошибка загрузки изображения: ${img.src}`);
// 		};
// 	});

// }());