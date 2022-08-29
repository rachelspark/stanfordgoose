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
    <div class="text-xs">
      {@html data.courseDescription
        .replaceAll("&nbsp;", "\xa0")
        .replaceAll(/<p>\s*<\/p>/g, "")}
    </div>
  </div>