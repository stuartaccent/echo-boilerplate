import htmx from 'htmx.org';

// Expose htmx globally (handy for debugging and extensions)
window.htmx = htmx;

// Normalize some HTTP error behaviors
document.body.addEventListener('htmx:beforeSwap', (evt) => {
    switch (evt.detail.xhr.status) {
        case 404:
            alert("Error: Not Found (404)");
            break;
        case 422:
            // Allow swapping validation errors into the page without marking as failure
            evt.detail.shouldSwap = true;
            evt.detail.isError = false;
            break;
        default:
            break;
    }
});

// Initialize Alpine.js components for any content injected by HTMX swaps.
// This ensures new fragments with x-data, x-init, etc., are activated.
function initAlpineIn(el) {
    if (window.Alpine && typeof window.Alpine.initTree === 'function' && el) {
        window.Alpine.initTree(el);
    }
}

// Run for any new content loaded by htmx (covers most swap scenarios)
htmx.onLoad((el) => {
    initAlpineIn(el);
});

// Fallback: after any swap on the target, (useful for some extensions/morph cases)
document.body.addEventListener('htmx:afterSwap', (evt) => {
    // evt.target is the element that was the target of the swap
    initAlpineIn(evt.target);
});