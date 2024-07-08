document.addEventListener("alpine:init", () => {
  Alpine.data("gloapp", () => ({
    editorTab: "icon", // icon, background
    previewSize: 0,
    previewSizeStyle: "",

    iconColor: "#ffffff",
    iconSize: "90",
    iconRotate: "0",
    bgColor: "#000000",
    bgRadius: "10",

    originalSVG: `
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
    `,

    createBackground(bgColor, bgRadius) {
      return `
        <rect
          width="100%"
          height="100%"
          fill="${bgColor}"
          rx="${bgRadius}%"
        ></rect>
      `
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
      const icon = this.createSvgIcon(this.iconColor, this.iconSize, this.iconRotate)
      const bg = this.createBackground(this.bgColor, this.bgRadius)
      return this.createParentSvg(size, bg + icon)
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
      return this.createSvg(this.previewSize)
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

    get finalSvg() {
      return this.createSvg("500px")
    },

    get finalSvgWhite() {
      return this.createSvgWhite("500px")
    },

    get finalSvgBlack() {
      return this.createSvgBlack("500px")
    },

    startPreviewSizeCalc() {
      const el = document.getElementById("preview-container")
      const calc = () => {
        const width = el.offsetWidth
        const height = el.offsetHeight
        const size = Math.min(width, height) - 100
        this.previewSize = size
        this.previewSizeStyle = `width: ${size}px; height: ${size}px;`
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

    async svgToPngBlobUrl(svgContent) {
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
          }, 'image/png');
        };

        img.onerror = function () {
          reject(new Error('Error loading SVG image.'));
        };

        img.src = url;
      })
    },

    async downloadZippedLogos() {
      const zip = new JSZip();

      const addToZip = async (name, blobUrl) => {
        const response = await fetch(blobUrl);
        const blob = await response.blob();
        zip.file(name, blob);
      }

      const finalSvg = this.svgToBlobUrl(this.finalSvg)
      const finalSvgWhite = this.svgToBlobUrl(this.finalSvgWhite)
      const finalSvgBlack = this.svgToBlobUrl(this.finalSvgBlack)
      const finalPng = await this.svgToPngBlobUrl(this.finalSvg)
      const finalPngWhite = await this.svgToPngBlobUrl(this.finalSvgWhite)
      const finalPngBlack = await this.svgToPngBlobUrl(this.finalSvgBlack)

      await addToZip('logo.svg', finalSvg)
      await addToZip('logo-white.svg', finalSvgWhite)
      await addToZip('logo-black.svg', finalSvgBlack)
      await addToZip('logo.png', finalPng)
      await addToZip('logo-white.png', finalPngWhite)
      await addToZip('logo-black.png', finalPngBlack)

      zip.generateAsync({ type: 'blob' }).then(function (content) {
        const link = document.createElement('a');
        link.href = URL.createObjectURL(content);
        link.download = `logo-${Date.now()}.zip`;
        link.click();

        URL.revokeObjectURL(link.href);
      });
    },

    init() {
      this.startPreviewSizeCalc()
      console.log("✨ Generate Logo Online initialized!")
    }
  }))
})