document.addEventListener('alpine:init', () => {
  Alpine.data('gloapp', () => ({
    init() {
      console.log('✨ Generate Logo Online initialized!')
    }
  }))
})