(function () {
  const inputEl = document.getElementById('input-text');
  const statusEl = document.getElementById('status-message');
  const clearBtn = document.getElementById('clear-btn');

  const outputs = {
    camel: document.getElementById('camel-output'),
    pascal: document.getElementById('pascal-output'),
    snake: document.getElementById('snake-output'),
    kebab: document.getElementById('kebab-output'),
    constant: document.getElementById('constant-output'),
    sentence: document.getElementById('sentence-output'),
    title: document.getElementById('title-output')
  };

  function setStatus(message) {
    statusEl.textContent = message;
  }

  function toWords(text) {
    const normalized = (text || '')
      .replace(/([a-z0-9])([A-Z])/g, '$1 $2')
      .replace(/[_\-]+/g, ' ')
      .replace(/[^\w\s]+/g, ' ')
      .trim()
      .toLowerCase();
    if (!normalized) {
      return [];
    }
    return normalized.split(/\s+/).filter(Boolean);
  }

  function toCamel(words) {
    if (words.length === 0) return '';
    return words[0] + words.slice(1).map(capitalize).join('');
  }

  function toPascal(words) {
    return words.map(capitalize).join('');
  }

  function toSnake(words) {
    return words.join('_');
  }

  function toKebab(words) {
    return words.join('-');
  }

  function toConstant(words) {
    return words.join('_').toUpperCase();
  }

  function toSentence(words) {
    if (words.length === 0) return '';
    const sentence = words.join(' ');
    return sentence.charAt(0).toUpperCase() + sentence.slice(1);
  }

  function toTitle(words) {
    return words.map(capitalize).join(' ');
  }

  function capitalize(word) {
    return word.charAt(0).toUpperCase() + word.slice(1);
  }

  function render() {
    const words = toWords(inputEl.value);

    outputs.camel.value = toCamel(words);
    outputs.pascal.value = toPascal(words);
    outputs.snake.value = toSnake(words);
    outputs.kebab.value = toKebab(words);
    outputs.constant.value = toConstant(words);
    outputs.sentence.value = toSentence(words);
    outputs.title.value = toTitle(words);

    if (words.length === 0) {
      setStatus('Enter text to convert.');
    } else {
      setStatus('Converted ' + words.length + ' words across all case styles.');
    }
  }

  async function copyFromTarget(targetId) {
    const element = document.getElementById(targetId);
    if (!element || !element.value) {
      setStatus('Nothing to copy for selected output.');
      return;
    }
    try {
      await navigator.clipboard.writeText(element.value);
      setStatus('Copied ' + targetId + ' to clipboard.');
    } catch (error) {
      setStatus('Clipboard copy failed. Please copy manually.');
    }
  }

  clearBtn.addEventListener('click', function () {
    inputEl.value = '';
    Object.values(outputs).forEach(function (el) { el.value = ''; });
    setStatus('Cleared.');
  });

  document.querySelectorAll('[data-copy-target]').forEach(function (button) {
    button.addEventListener('click', function () {
      copyFromTarget(button.getAttribute('data-copy-target'));
    });
  });

  inputEl.addEventListener('input', render);
  render();
})();