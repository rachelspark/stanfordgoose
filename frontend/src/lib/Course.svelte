<script lang="ts">
import type { CourseData } from "./search";

    import Tag from "./Tag.svelte"
    export let data: CourseData;

    const TagColors = {
      "Aut": "pink",
      "Win": "lightblue",
      "Spr": "lightyellow",
      "Sum": "palegreen",
      "lastOffered" : "lightgray"
    }

  </script>
  
  <div>
    <div class="flex flex-row">
    <h3 class="basis-3/4 text-sm font-bold">
      {data.dept}
      {data.courseNumber}:
      {data.courseTitle}
    </h3>
    <div class="basis-1/4">
    <h6 class="flex flex-row justify-end text-xs font-light">
      {#if data.lastOffered == ""}
        {#if data.terms}
        {#each data.terms as term, i}
          <Tag data={term} color={TagColors[term]}/>
        {/each}
        {/if}
        {:else}
        <Tag data={data.lastOffered.split("| ")[0]} color={TagColors["lastOffered"]}/>
      {/if}
    </h6>
      </div>
  </div>
    <p class="prof-links">
      {#if data.instructors}
      {#each data.instructors as instructor, i}
      <a target="_blank"
        href= {instructor.profileUrl}>{instructor.name}</a
      >{#if i < data.instructors.length - 1}{", "}{/if}
      {/each}
      {/if}
    </p>
    <p class="text-xs font-light mb-1">
      {#if data.lastOffered == "" && data.units != ""}
      {(data.units == "1") ? `${data.units} unit`: `${data.units} units`} |
        {data.ugReqs ?         `${data.ugReqs.join(", ")}` : "No UG Reqs"}
      {/if}
      
    </p>
    <div class="text-xs mb-1">
      {@html data.courseDescription
        .replaceAll("&nbsp;", "\xa0")
        .replaceAll(/<p>\s*<\/p>/g, "")}
    </div>
    <div class="ext-links flex space-x-2">
      <a
        target="_blank"
        rel="noopener noreferrer"
        href={(data.dept && data.courseNumber)
          ? `https://carta-beta.stanford.edu/course/${data.dept}${data.courseNumber}`
          : `https://carta-beta.stanford.edu/results/${data.dept.toLowerCase}%20${data.courseNumber}`}
        >Carta</a
      >
      <a
        target="_blank"
        rel="noopener noreferrer"
        href={(data.dept && data.courseNumber)
          ? `https://explorecourses.stanford.edu/search?view=catalog&filter-coursestatus-Active=on&page=0&catalog=&academicYear=&q=${data.dept}+${data.courseNumber}`
          : `https://explorecourses.stanford.edu/`}
        >ExploreCourses</a
      >
    </div>
  </div>

  <style lang="postcss">
    .ext-links a {
      @apply text-gray-500 text-sm hover:underline;
    }
    .prof-links a {
      @apply text-black text-sm mb-1 hover:underline;
    }
  </style>