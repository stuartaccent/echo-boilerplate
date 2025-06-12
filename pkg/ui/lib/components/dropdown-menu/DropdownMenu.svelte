<svelte:options customElement={{
    tag: "ui-dropdown-menu",
    shadow: "none",
    props: {
        items: { type: "Array" }
    }
}}/>

<script lang="ts">
  import {onMount} from "svelte";
  import {fade} from "svelte/transition";

  type MenuItemType = "label" | "separator" | "item";

  const MENU_ITEM_TYPE = {
    LABEL: "label" as MenuItemType,
    SEPARATOR: "separator" as MenuItemType,
    ITEM: "item" as MenuItemType
  };

  interface MenuItem {
    type: MenuItemType;
    text?: string;
    href?: string;
  }

  interface MenuProps {
    title: string;
    items: MenuItem[];
    position: string;
  }

  let {
    title = "Menu",
    items = [],
    position = "left-0 top-[110%]"
  }: MenuProps = $props();

  let isOpen = $state(false);
  let ref = $state<HTMLDivElement | null>(null);

  function clickOutside(event: MouseEvent): void {
    if (isOpen && ref && !ref.contains(event.target as Node)) {
      isOpen = false;
    }
  }

  onMount(() => {
    document.addEventListener("click", clickOutside);
    return () => {
      document.removeEventListener("click", clickOutside);
    };
  });

  function toggle(): void {
    isOpen = !isOpen;
  }
</script>

<div class="owl-dropdown-menu" bind:this={ref}>
  <button onclick={toggle} class="owl-button owl-button-ghost">
    <span>{title}</span>
    <ui-icon icon="chevron-down"></ui-icon>
  </button>
  {#if isOpen && items?.length > 0}
    <div class="owl-dropdown-menu-content {position}" role="menu" transition:fade={{ duration: 200 }}>
      {#each items as item}
        {#if item.type === MENU_ITEM_TYPE.LABEL}
          <div class="owl-dropdown-menu-label">{item.text}</div>
        {:else if item.type === MENU_ITEM_TYPE.SEPARATOR}
          <div class="owl-dropdown-menu-separator" role="presentation"></div>
        {:else if item.type === MENU_ITEM_TYPE.ITEM}
          <a class="owl-dropdown-menu-item" href={item.href} role="menuitem">{item.text}</a>
        {/if}
      {/each}
    </div>
  {/if}
</div>
