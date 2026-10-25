(() => {
  document.querySelectorAll('.quiz').forEach((quiz) => {
    const buttons = [...quiz.querySelectorAll('.quiz-option')];
    const feedback = quiz.querySelector('.quiz-feedback');
    const reset = quiz.querySelector('.quiz-reset');

    const clear = () => {
      buttons.forEach((button) => {
        button.disabled = false;
        button.classList.remove('is-correct', 'is-wrong');
      });
      feedback.textContent = '';
    };

    buttons.forEach((button) => {
      button.addEventListener('click', () => {
        const correct = button.dataset.correct === 'true';
        buttons.forEach((option) => {
          option.disabled = true;
          if (option.dataset.correct === 'true') option.classList.add('is-correct');
        });
        if (!correct) button.classList.add('is-wrong');
        feedback.textContent = `${correct ? 'Correct.' : 'Not quite.'} ${quiz.dataset.explanation}`;
      });
    });

    reset?.addEventListener('click', clear);
  });
})();
