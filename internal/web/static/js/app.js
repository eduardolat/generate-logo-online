document.addEventListener("alpine:init", () => {
  Alpine.data("gloapp", () => ({
    editorTab: "icon", // icon, background
    previewSize: 0,
    previewSizeStyle: "",

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