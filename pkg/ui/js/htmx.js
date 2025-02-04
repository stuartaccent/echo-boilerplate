import htmx from 'htmx.org';

window.htmx = htmx;

document.addEventListener('DOMContentLoaded', () => {
    document.body.addEventListener('htmx:beforeSwap', (evt) => {
        switch (evt.detail.xhr.status) {
            case 404:
                alert("Error: Not Found (404)");
                break;
            case 422:
                evt.detail.shouldSwap = true;
                evt.detail.isError = false;
                break;
            default:
                break;
        }
    });
});