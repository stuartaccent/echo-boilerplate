<svelte:options customElement={{
  tag: "ui-icon",
  shadow: "open"
}}/>

<script lang="ts">
  const {icon = "", width = "1em", height = "1em"} = $props<{
    icon?: string;
    width?: string;
    height?: string;
  }>();
</script>

<svg xmlns="http://www.w3.org/2000/svg" {width} {height}>
  <use xlink:href="/static/svg/icons.svg#icon-{icon}"></use>
</svg>
