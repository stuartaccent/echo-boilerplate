import Alpine from 'alpinejs';
import anchor from '@alpinejs/anchor';

Alpine.plugin(anchor)

Alpine.data('menu', () => ({
    open: false,

    toggle() {
        this.open = ! this.open
    },
    close() {
        this.open = false
    },
}));

window.Alpine = Alpine
Alpine.start();
