<script lang="ts">
    import type { CourseData } from "./search";
  
    export let data: CourseData;
  </script>
  
  <div>
    <h3 class="text-sm font-bold">
      {data.courseDept}
      {data.courseNumber}
      {data.courseTitle}
    </h3>
    <!-- <p class="text-sm mb-1">
        {#each data.instructors as instructor, i}
          <a href={instructor.profileURL}
            >{instructor.name}</a
          >{#if i < data.instructors.length - 1}{", "}{/if}
        {/each}
      </p>
      <p class="text-xs font-light mb-1">
        {data.terms} | {data.units} | {data.ugReqs}
      </p> -->
    <div class="text-xs">
      {@html data.courseDescription
        .replaceAll("&nbsp;", "\xa0")
        .replaceAll(/<p>\s*<\/p>/g, "")}
    </div>
  </div>