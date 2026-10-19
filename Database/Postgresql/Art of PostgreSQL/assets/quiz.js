document.querySelectorAll(".quiz").forEach((quiz) => {
  const feedback = quiz.querySelector(".quiz-feedback");
  const buttons = [...quiz.querySelectorAll(".quiz-option")];
  const correct = quiz.dataset.correct;

  buttons.forEach((button, index) => {
    button.addEventListener("click", () => {
      const isCorrect = button.dataset.option === correct;
      buttons.forEach((option) => {
        option.disabled = true;
        if (option.dataset.option === correct) option.classList.add("correct");
      });
      if (!isCorrect) button.classList.add("incorrect");
      feedback.textContent = isCorrect
        ? `Correct. ${quiz.dataset.explanation}`
        : `Not quite. ${quiz.dataset.explanation}`;
      feedback.setAttribute("role", "status");
      feedback.setAttribute("aria-live", "polite");
    }, { once: true });
  });
});
