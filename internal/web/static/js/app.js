document.addEventListener("alpine:init", () => {
  Alpine.data("gloapp", () => ({
    isMobileEditorOpen: false,
    editorTab: "icon", // icon, background

    init() {
      console.log("✨ Generate Logo Online initialized!")
    }
  }))
})