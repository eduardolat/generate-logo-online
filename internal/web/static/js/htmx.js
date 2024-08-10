export function initHTMX () {
  document.addEventListener('DOMContentLoaded', () => {
    /*
    This fixes this issue:
    https://stackoverflow.com/questions/73658449/htmx-request-not-firing-when-hx-attributes-are-added-dynamically-from-javascrip
  */
    const observer = new MutationObserver((mutations) => {
      mutations.forEach((mutation) => {
        mutation.addedNodes.forEach((node) => {
          if (node.nodeType === 1 && !node['htmx-internal-data']) {
            htmx.process(node)
          }
        })
      })
    })
    observer.observe(document, { childList: true, subtree: true })
  })
}
