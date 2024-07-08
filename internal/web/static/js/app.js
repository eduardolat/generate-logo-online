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

    init() {
      this.startPreviewSizeCalc()
      console.log("✨ Generate Logo Online initialized!")
    }
  }))
})