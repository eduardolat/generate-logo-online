document.addEventListener("alpine:init", () => {
  Alpine.data("gloapp", () => ({
    editorTab: "icon", // icon, background
    previewSize: 0,

    iconColor: "#ffffff",
    iconSize: "90",
    iconRotate: "0",

    bgType: "solid", // solid, gradient
    bgRadius: "10",
    bgColor: "#000000",
    bgGradientStart: "#000000",
    bgGradientEnd: "#ffffff",
    bgGradientType: "linear", // linear, radial
    bgGradientAngle: "0",
    bgGradientCutLine: "50",
    bgGradientBlur: "50",

    originalSVG: "",

    setDefaultValues() {
      this.editorTab = "icon"

      this.iconColor = "#ffffff"
      this.iconSize = "90"
      this.iconRotate = "0"

      this.bgType = "solid"
      this.bgRadius = "10"
      this.bgColor = "#000000"
      this.bgGradientStart = "#000000"
      this.bgGradientEnd = "#001475"
      this.bgGradientType = "linear"
      this.bgGradientAngle = "0"
      this.bgGradientCutLine = "50"
      this.bgGradientBlur = "50"

      this.originalSVG = `
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M9.937 15.5A2 2 0 0 0 8.5 14.063l-6.135-1.582a.5.5 0 0 1 0-.962L8.5 9.936A2 2 0 0 0 9.937 8.5l1.582-6.135a.5.5 0 0 1 .963 0L14.063 8.5A2 2 0 0 0 15.5 9.937l6.135 1.581a.5.5 0 0 1 0 .964L15.5 14.063a2 2 0 0 0-1.437 1.437l-1.582 6.135a.5.5 0 0 1-.963 0z" />
          <path d="M20 3v4" />
          <path d="M22 5h-4" />
          <path d="M4 17v2" />
          <path d="M5 18H3" />
        </svg>
      `
    },

    setOriginalSVG(svg) {
      this.originalSVG = svg

      // Close the modal simulating the user pressing the Escape key
      document.dispatchEvent(new KeyboardEvent('keyup', {
        bubbles: true,
        cancelable: true,
        key: 'Escape',
        code: 'Escape',
        keyCode: 27,
        which: 27
      }))
    },

    createBackground(
      bgType, bgColor, bgRadius, bgGradientStart, bgGradientEnd,
      bgGradientType, bgGradientAngle, bgGradientCutLine, bgGradientBlur
    ) {
      if (bgType === "solid") {
        return `
          <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%">
            <rect
              width="100%"
              height="100%"
              fill="${bgColor}"
              rx="${bgRadius}%"
            ></rect>
          </svg>
        `;
      }

      if (bgType === "gradient") {
        const gradID = `grad-bg-${crypto.randomUUID()}`;
        let gradient;

        // Normalize bgGradientCutLine and bgGradientBlur to values between 0 and 1
        const cutLine = bgGradientCutLine / 100;
        const blur = bgGradientBlur / 100;

        // Calculate the stop points for the gradient
        const startStop = Math.max(0, cutLine - blur / 2);
        const endStop = Math.min(1, cutLine + blur / 2);

        if (bgGradientType === "linear") {
          const angleRad = (bgGradientAngle - 90) * (Math.PI / 180);
          const x1 = 50 + Math.cos(angleRad) * 50;
          const y1 = 50 + Math.sin(angleRad) * 50;
          const x2 = 50 - Math.cos(angleRad) * 50;
          const y2 = 50 - Math.sin(angleRad) * 50;

          gradient = `
            <linearGradient id="${gradID}" x1="${x1}%" y1="${y1}%" x2="${x2}%" y2="${y2}%">
              <stop offset="0%" stop-color="${bgGradientStart}" />
              <stop offset="${startStop * 100}%" stop-color="${bgGradientStart}" />
              <stop offset="${endStop * 100}%" stop-color="${bgGradientEnd}" />
              <stop offset="100%" stop-color="${bgGradientEnd}" />
            </linearGradient>
          `;
        }

        if (bgGradientType === "radial") {
          gradient = `
            <radialGradient id="${gradID}" cx="50%" cy="50%" r="50%" fx="50%" fy="50%">
              <stop offset="0%" stop-color="${bgGradientStart}" />
              <stop offset="${startStop * 100}%" stop-color="${bgGradientStart}" />
              <stop offset="${endStop * 100}%" stop-color="${bgGradientEnd}" />
              <stop offset="100%" stop-color="${bgGradientEnd}" />
            </radialGradient>
          `;
        }

        return `
          <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%">
            <defs>${gradient}</defs>
            <rect width="100%" height="100%" fill="url(#${gradID})" rx="${bgRadius}%" />
          </svg>
        `;
      }

      return "";
    },

    createSvgIcon(iconColor, iconSize, iconRotate) {
      const svg = this.originalSVG
        .replace(/stroke="currentColor"/g, `stroke="${iconColor}"`)

      const scale = iconSize * 0.01
      return `
        <g transform="translate(12, 12) rotate(${iconRotate}) scale(${scale}) translate(-12, -12)">
          ${svg}
        </g>
      `
    },

    createParentSvg(size, children) {
      if (size <= 0) size = 0
      return `
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="${size}"
          height="${size}"
          viewBox="0 0 24 24"
        >
          ${children}
        </svg>
      `
    },

    createSvg(size) {
      const icon = this.createSvgIcon(this.iconColor, this.iconSize, this.iconRotate);
      const bg = this.createBackground(
        this.bgType,
        this.bgColor,
        this.bgRadius,
        this.bgGradientStart,
        this.bgGradientEnd,
        this.bgGradientType,
        this.bgGradientAngle,
        this.bgGradientCutLine,
        this.bgGradientBlur
      );
      return this.createParentSvg(size, bg + icon);
    },

    createSvgWhite(size) {
      const icon = this.createSvgIcon("#ffffff", this.iconSize, this.iconRotate)
      return this.createParentSvg(size, icon)
    },

    createSvgBlack(size) {
      const icon = this.createSvgIcon("#000000", this.iconSize, this.iconRotate)
      return this.createParentSvg(size, icon)
    },

    get previewSvg() {
      return this.createSvg(this.previewSize);
    },

    get dlPreviewSvg() {
      return this.createSvg("150px")
    },

    get dlPreviewSvgWhite() {
      return this.createSvgWhite("150px")
    },

    get dlPreviewSvgBlack() {
      return this.createSvgBlack("150px")
    },

    startPreviewSizeCalc() {
      const el = document.getElementById("preview-container")
      const calc = () => {
        const width = el.offsetWidth
        const height = el.offsetHeight
        const size = Math.min(width, height) - 100
        this.previewSize = size
      }

      if (el) {
        calc()
        const resizeObserver = new ResizeObserver(() => calc())
        resizeObserver.observe(el)
      }
    },

    svgToBlobUrl(svgContent) {
      const svgBlob = new Blob([svgContent], { type: 'image/svg+xml' })
      return URL.createObjectURL(svgBlob)
    },

    async svgToImageBlobUrl(svgContent, mimeType) {
      const url = this.svgToBlobUrl(svgContent)

      return new Promise((resolve, reject) => {
        const img = new Image();

        img.onload = function () {
          URL.revokeObjectURL(url);

          const canvas = document.createElement('canvas');
          canvas.width = img.width;
          canvas.height = img.height;

          const ctx = canvas.getContext('2d');
          ctx.drawImage(img, 0, 0);

          canvas.toBlob(function (blob) {
            if (blob) {
              const pngUrl = URL.createObjectURL(blob);
              resolve(pngUrl);
            } else {
              reject(new Error('Error creating PNG blob.'));
            }

            canvas.remove();
          }, mimeType);
        };

        img.onerror = function () {
          reject(new Error('Error loading SVG image.'));
        };

        img.src = url;
      })
    },

    async svgToPngBlobUrl(svgContent) {
      return this.svgToImageBlobUrl(svgContent, 'image/png')
    },

    async svgToIcoBlobUrl(svgContent) {
      return this.svgToImageBlobUrl(svgContent, 'image/x-icon')
    },

    async downloadZippedLogos() {
      const zip = new JSZip();

      const addToZip = async (name, blobUrl) => {
        const response = await fetch(blobUrl);
        const blob = await response.blob();
        zip.file(name, blob);
      }

      const filesToZip = [
        { name: "svg/logo.svg", size: "512px", color: "default" },
        { name: "svg/logo-white.svg", size: "512px", color: "white" },
        { name: "svg/logo-black.svg", size: "512px", color: "black" },

        { name: "png/logo-512.png", size: "512px", color: "default" },
        { name: "png/logo-256.png", size: "256px", color: "default" },
        { name: "png/logo-128.png", size: "128px", color: "default" },
        { name: "png/logo-64.png", size: "64px", color: "default" },
        { name: "png/logo-32.png", size: "32px", color: "default" },
        { name: "png/logo-16.png", size: "16px", color: "default" },
        { name: "png/logo-white-512.png", size: "512px", color: "white" },
        { name: "png/logo-white-256.png", size: "256px", color: "white" },
        { name: "png/logo-white-128.png", size: "128px", color: "white" },
        { name: "png/logo-white-64.png", size: "64px", color: "white" },
        { name: "png/logo-white-32.png", size: "32px", color: "white" },
        { name: "png/logo-white-16.png", size: "16px", color: "white" },
        { name: "png/logo-black-512.png", size: "512px", color: "black" },
        { name: "png/logo-black-256.png", size: "256px", color: "black" },
        { name: "png/logo-black-128.png", size: "128px", color: "black" },
        { name: "png/logo-black-64.png", size: "64px", color: "black" },
        { name: "png/logo-black-32.png", size: "32px", color: "black" },
        { name: "png/logo-black-16.png", size: "16px", color: "black" },

        { name: "ico/logo-64.ico", size: "64px", color: "default" },
        { name: "ico/logo-32.ico", size: "32px", color: "default" },
        { name: "ico/logo-16.ico", size: "16px", color: "default" },
        { name: "ico/logo-white-64.ico", size: "64px", color: "white" },
        { name: "ico/logo-white-32.ico", size: "32px", color: "white" },
        { name: "ico/logo-white-16.ico", size: "16px", color: "white" },
        { name: "ico/logo-black-64.ico", size: "64px", color: "black" },
        { name: "ico/logo-black-32.ico", size: "32px", color: "black" },
        { name: "ico/logo-black-16.ico", size: "16px", color: "black" }
      ]

      const getSvg = (size, color) => {
        if (color === "default") return this.createSvg(size);
        if (color === "white") return this.createSvgWhite(size);
        if (color === "black") return this.createSvgBlack(size);
      }

      for await (const file of filesToZip) {
        const svgStrig = getSvg(file.size, file.color)

        if (file.name.includes(".svg")) {
          const svg = this.svgToBlobUrl(svgStrig)
          await addToZip(file.name, svg)
        }

        if (file.name.includes(".png")) {
          const png = await this.svgToPngBlobUrl(svgStrig)
          await addToZip(file.name, png)
        }

        if (file.name.includes(".ico")) {
          const ico = await this.svgToIcoBlobUrl(svgStrig)
          await addToZip(file.name, ico)
        }
      }

      zip.generateAsync({ type: 'blob' }).then(function (content) {
        const link = document.createElement('a');
        link.href = URL.createObjectURL(content);
        link.download = `logo-${Date.now()}.zip`;
        link.click();

        URL.revokeObjectURL(link.href);
        link.remove();
      });
    },

    init() {
      this.startPreviewSizeCalc()
      this.setDefaultValues()
      console.log("✨ Generate Logo Online initialized!")
      console.log("Star on GitHub: https://github.com/eduardolat/generate-logo")
    }
  }))
})
