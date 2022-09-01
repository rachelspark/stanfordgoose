<script lang="ts" context="module">
  export function encodeQueryHash(query: string): string {
    return "#" + encodeURIComponent(query).replaceAll("%20", "+");
  }

  export function decodeQueryHash(hash: string): string {
    return decodeURIComponent(hash.slice(1).replaceAll("+", "%20"));
  }
</script>

<script lang="ts">
  import { onMount } from "svelte";
  import gooseLogo from "./assets/goose-logo-gradient.svg"
  import Course from "./lib/Course.svelte";
  import QueryLink from "./lib/QueryLink.svelte";
  import Footer from "./lib/Footer.svelte"
  import { createSearcher, normalizeText } from "./lib/search";

  let learnMore;

  let query: string = location.hash ? decodeQueryHash(location.hash) : "";
  $: {
    const newUrl = query
      ? encodeQueryHash(query)
      : location.pathname + location.search;
    history.replaceState(null, "", newUrl);
  }

  let landing = query === "";
  $: if (query) landing = false;

  let activeCourses = false;

  const { data, error, search } = createSearcher();
  $: finalQuery = (activeCourses ? "@terms:{Aut | Win | Spr | Sum}" : "") + normalizeText(query);
  $: search(finalQuery);

  let learnMoreShown = false;

  // Render courses incrementally in batches of 20 at a time, to avoid slowing
  // down the browser with too many elements at once.
  let showing = 0;
  let showingTimeout = 0;
  function showMore() {
    const len = $data?.courses?.length ?? 0;
    if (showing < len) {
      showing += Math.min(20, len - showing);
      showingTimeout = window.setTimeout(showMore, 100);
    }
  }
  onMount(() =>
    data.subscribe(() => {
      window.clearTimeout(showingTimeout);
      showing = 0;
      showMore();
    })
  );
</script>

<main class="px-4 py-8 max-w-screen-md mx-auto" class:landing>
  <div class="flex flex-row justify-center items-center">
      <a
        href="/"
        on:click|preventDefault={() => ((query = ""), (landing = true))}
        ><img src={gooseLogo} alt="GOOSE" class="container max-w-36 max-h-14 top-0 pb-4"/></a
      >
  </div>
    {#if !landing}
      <div class="space-y-2 mb-3 text-sm text-zinc-600">
        <p>
          Try words, phrases, titles, subjects, course numbers, and instructor last
          names. You can also look for exact textual phrases (like
          <QueryLink bind:query value={`"creative process"`} />) and prefix
          matches (such as
          <QueryLink bind:query value={`genom*`} />).
        </p>
        <p>
          Filter by specific attributes like
          <QueryLink bind:query value={`@dept:"CS"`} />, 
          <QueryLink bind:query value={`@terms:{Aut}`} />,
          <QueryLink bind:query value={`@ugReqs:{WAY-EDP}`} />, and
          <QueryLink bind:query value={`@level:{Intro}`} />.
        </p>
        <p>
          You can also combine filters. For example, let's say this fall or spring I want to fulfill my WAY-ER or WAY-A-II with a course about media:
          <QueryLink bind:query value={`@terms:{Aut | Spr} @ugReqs:{WAY-A-II | WAY-ER} media`} />
        </p>
      </div>
    {/if}
    <div class="flex flex-col items-center">
    <h1 class="font-bold">
      <span class="flavor">The fastest way to search Stanford courses.</span>
    </h1>
    <p class="container max-w-lg">
      <span class="flavor">Consider the wild-goose chase of course selection over. Goose gives you more useful results, crazy fast.</span>
    </p>

      <div class="relative searchbar-wrapper z-10"
        >
        <div class="flex absolute inset-y-0 left-3 items-center">
          <svg
            class="w-5 h-5 text-gray-400 pointer-events-none"
            fill="currentColor"
            viewBox="0 0 50 50"
            ><path
              d="M 21 3 C 11.601563 3 4 10.601563 4 20 C 4 29.398438 11.601563 37 21 37 C 24.355469 37 27.460938 36.015625 30.09375 34.34375 L 42.375 46.625 L 46.625 42.375 L 34.5 30.28125 C 36.679688 27.421875 38 23.878906 38 20 C 38 10.601563 30.398438 3 21 3 Z M 21 7 C 28.199219 7 34 12.800781 34 20 C 34 27.199219 28.199219 33 21 33 C 13.800781 33 8 27.199219 8 20 C 8 12.800781 13.800781 7 21 7 Z"
            />
          </svg
          >
          </div>
          <!-- svelte-ignore a11y-autofocus -->
          <input
          autofocus
          class="searchbar w-full text-md pl-10 rounded-full bg-gray-50 hover:bg-gray-100 focus:outline-none border"
          placeholder={landing ? "I want to take a class in..." : "Search..."}
          bind:value={query}
        /></div
      >
        </div>

    {#if !landing}
    <label class="flex text-sm mb-2 ml-2 mt-2">
      <input class="mr-2" type="checkbox" bind:checked={activeCourses} />
      Only show courses offered in 2022-23
    </label>
  {/if}

  {#if $error !== null}
    <p class="text-red-500 mb-4">
      {$error}
    </p>
  {/if}
  {#if query && $data}
    <p class="text-sm mb-4 bg-red-50 px-2 py-1 text-rose-600">
      Found {$data.count} results
      <span class="text-rose-400">({($data.time * 1000).toFixed(2)} ms)</span>
    </p>
    {#if $data.count == 0}Silly goose! {/if}

    <div class="space-y-4">
      {#each ($data.courses ?? []).slice(0, showing) as course (course.id)}
        <Course data={course} />
      {/each}
    </div>
  {/if}

  {#if landing}
  <div class="flex flex-col items-center text-tiny">
  <div class="flex flex-row text-tiny pt-1 gap-x-4">
    <button class="text-buttons" on:click={() => (learnMoreShown = !learnMoreShown)}>Learn More</button>
    <a target="_blank" class="text-buttons" href="https://3jysmfizo7l.typeform.com/to/BQ5rLXna">Send Feedback</a>
  </div>
  {#if learnMoreShown}        
  <div class="box container max-w-sm text-xs text-center text-gray-300 mt-2">Built in Go and Svelte on an in-memory Redis database, deployed to the edge using Fly.io <br><br>  Built with ❤️ by a fellow Cardinal frustrated by ExploreCourses <br> This project also uses some code ported from classes.wtf and licensed by its author: Copyright 2022 Eric Zhang </div>
{/if}
</div>
  <div class="absolute bottom-0 z-10"><Footer/></div>
  {/if}
</main>

<style lang="postcss">
    .landing {
      @apply flex flex-col justify-center;
    }

    .landing h1 {
      @apply text-center text-xl text-black mb-4 pt-6;
    }

    .landing p {
      @apply text-base text-center text-black mb-6;
    }
    
    .landing input {
      @apply min-w-[20ch] text-sm px-3 py-2 pl-10;;
    }

    .text-buttons {
      @apply text-xs text-gray-300 underline hover:text-gray-500;
    }

  @screen md {
    .landing {
      @apply flex flex-col justify-center space-y-4;
    }

    .landing h1 {
      @apply text-center text-6xl text-black mb-8 pt-14 pb-2;
    }

    .landing p {
      @apply text-xl text-center text-black mb-8 pb-8;
    }
    
    .landing input {
      @apply w-[50ch] text-base px-3 py-2 pl-10;;
    }

    .text-buttons {
      @apply text-gray-300 underline hover:text-gray-600;
    }
  }

  main:not(.landing) .flavor {
    @apply hidden;
  }

  main:not(.landing) .searchbar-wrapper {
    @apply text-base w-full;
  }

  main:not(.landing) .searchbar {
    @apply w-full px-3 py-2 pl-10;
  }
</style>
