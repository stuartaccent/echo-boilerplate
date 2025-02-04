class OwlIcon extends HTMLElement {
    constructor() {
        super();

        const shadow = this.attachShadow({ mode: "open" });

        this.svgElement = document.createElementNS("http://www.w3.org/2000/svg", "svg");
        this.svgElement.setAttribute("xmlns", "http://www.w3.org/2000/svg");
        this.useElement = document.createElementNS("http://www.w3.org/2000/svg", "use");
        this.svgElement.appendChild(this.useElement);

        shadow.appendChild(this.svgElement);

        this.svgElement.setAttribute("width", "1em");
        this.svgElement.setAttribute("height", "1em");
    }

    static get observedAttributes() {
        return ["icon", "width", "height"];
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (oldValue !== newValue) {
            this.updateAttributes(name);
        }
    }

    updateAttributes(attribute) {
        const icon = this.getAttribute("icon");
        const width = this.getAttribute("width");
        const height = this.getAttribute("height");
        switch (attribute) {
            case "icon":
                if (icon) {
                    this.useElement.setAttributeNS(
                        "http://www.w3.org/1999/xlink",
                        "href",
                        `static/svg/icons.svg#icon-${icon}`
                    );
                } else {
                    console.warn("Missing or invalid icon attribute");
                }
                break;
            case "width":
            case "height":
                this.svgElement.setAttribute(attribute, attribute === "width" ? width : height);
                break;
        }
    }
}

customElements.define("owl-icon", OwlIcon);
