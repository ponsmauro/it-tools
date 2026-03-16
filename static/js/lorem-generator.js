(function initLoremGeneratorTool() {
        const wordPool = [
            "lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
            "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
            "magna", "aliqua", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation",
            "ullamco", "laboris", "nisi", "aliquip", "ex", "ea", "commodo", "consequat",
            "duis", "aute", "irure", "reprehenderit", "in", "voluptate", "velit", "esse",
            "cillum", "eu", "fugiat", "nulla", "pariatur", "excepteur", "sint", "occaecat",
            "cupidatat", "non", "proident", "sunt", "culpa", "qui", "officia", "deserunt",
            "mollit", "anim", "id", "est", "laborum"
        ];

        const typeSelect = document.getElementById("lorem-type");
        const countInput = document.getElementById("lorem-count");
        const output = document.getElementById("lorem-output");
        const status = document.getElementById("lorem-status");
        const generateBtn = document.getElementById("btn-lorem-generate");
        const copyBtn = document.getElementById("btn-lorem-copy");
        const clearBtn = document.getElementById("btn-lorem-clear");

        if (!typeSelect || !countInput || !output || !status || !generateBtn || !copyBtn || !clearBtn) {
            return;
        }

        function setStatus(message, isError) {
            status.textContent = message;
            status.classList.toggle("error", Boolean(isError));
        }

        function randomWord() {
            const index = Math.floor(Math.random() * wordPool.length);
            return wordPool[index];
        }

        function generateWords(count) {
            const words = [];
            for (let i = 0; i < count; i++) {
                words.push(randomWord());
            }
            const sentence = words.join(" ");
            return sentence.charAt(0).toUpperCase() + sentence.slice(1) + ".";
        }

        function generateParagraph() {
            const sentenceCount = 3 + Math.floor(Math.random() * 4);
            const sentences = [];
            for (let i = 0; i < sentenceCount; i++) {
                const wordsInSentence = 8 + Math.floor(Math.random() * 10);
                sentences.push(generateWords(wordsInSentence));
            }
            return sentences.join(" ");
        }

        function getSafeCount() {
            const value = Number.parseInt(countInput.value, 10);
            if (Number.isNaN(value)) {
                return 1;
            }
            return Math.min(100, Math.max(1, value));
        }

        generateBtn.addEventListener("click", function () {
            const mode = typeSelect.value;
            const count = getSafeCount();

            if (mode === "words") {
                output.value = generateWords(count);
                setStatus("Generated words.", false);
                return;
            }

            const paragraphs = [];
            for (let i = 0; i < count; i++) {
                paragraphs.push(generateParagraph());
            }
            output.value = paragraphs.join("\n\n");
            setStatus("Generated paragraphs.", false);
        });

        copyBtn.addEventListener("click", async function () {
            try {
                await navigator.clipboard.writeText(output.value);
                setStatus("Output copied.", false);
            } catch (error) {
                setStatus("Clipboard copy failed.", true);
            }
        });

        clearBtn.addEventListener("click", function () {
            output.value = "";
            setStatus("Output cleared.", false);
        });
    })();