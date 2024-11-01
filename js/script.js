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