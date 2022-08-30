<script lang="ts">
    import type { CourseData } from "./search";
  
    export let data: CourseData;
  </script>
  
  <div>
    <h3 class="text-sm font-bold">
      {data.courseDept}
      {data.courseNumber}:
      {data.courseTitle}
    </h3>
      <p class="text-xs font-light mb-1">
        {#if data.lastOffered == ""}
          {data.terms} | {(data.units == "1") ? `${data.units} unit`: `${data.units} units`} | {data.ugReqs ? data.ugReqs : "No UG reqs"}
        {:else}
          {data.lastOffered}
        {/if}
      </p>
    <div class="text-xs">
      {@html data.courseDescription
        .replaceAll("&nbsp;", "\xa0")
        .replaceAll(/<p>\s*<\/p>/g, "")}
    </div>
  </div>

  <div class="ext-links flex space-x-2">
    <a
      target="_blank"
      rel="noopener noreferrer"
      href={(data.courseDept && data.courseNumber)
        ? `https://carta-beta.stanford.edu/course/${data.courseDept}${data.courseNumber}`
        : `https://carta-beta.stanford.edu/results/${data.courseDept.toLowerCase}%20${data.courseNumber}`}
      >Carta</a
    >
    <a
      target="_blank"
      rel="noopener noreferrer"
      href={(data.courseDept && data.courseNumber)
        ? `https://explorecourses.stanford.edu/search?view=catalog&filter-coursestatus-Active=on&page=0&catalog=&academicYear=&q=${data.courseDept}+${data.courseNumber}`
        : `https://explorecourses.stanford.edu/`}
      >ExploreCourses</a
    >
  </div>

  <style lang="postcss">
    .ext-links a {
      @apply text-gray-500 text-sm underline hover:text-black;
    }
  </style>