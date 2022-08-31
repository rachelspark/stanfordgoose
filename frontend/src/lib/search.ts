import { writable, type Readable } from "svelte/store";

export type CourseData = {
  id: string;
  dept: string;
  deptLongname: string
  courseNumber: string;
  courseTitle: string;
  courseDescription: string;
  level: string;
  terms: string[];
  units: string;
  lastOffered: string
  instructors: Array<{
    name: string;
    isPI: string;
    profileUrl: string;
  }>;
  ugReqs: string[];
};

export type Searcher = {
    data: Readable<any>;
    error: Readable<string | null>;
    search: (query: string) => void;
};

export function createSearcher(): Searcher {
    let abort: AbortController | null = null;
    let lastQuery: string | null = null;
  
    const data = writable<any>(undefined);
    const error = writable<string | null>(null);
    const search = async (query: string) => {
      if (query === lastQuery) return;
      lastQuery = query;
      abort?.abort();
      abort = new AbortController();
      let localAbort = abort;
      try {
        const resp = await fetch("/search?q=" + encodeURIComponent(query), {
          signal: abort.signal,
        });
        if (!resp.ok) {
          const obj = await resp.json();
          error.set(obj.error);
        } else {
          const obj = await resp.json();
          data.set(obj);
          error.set(null);
        }
      } catch (err) {
        if (!localAbort.signal.aborted) {
          // Network error or some other issue.
          error.set(err.message);
        }
      }
    };
  
    return { data, error, search };
  }
  

/** Apply some transformations to a query to make it more useful by default. */
export function normalizeText(query: string): string {
  if (query.length >= 2 && query.slice(-2).match(/\w{2}/)) {
    const i = /\w+$/.exec(query).index;
    const partial = query.substring(i);
    query = query.substring(0, i) + `(${partial}|${partial}*)`; // prefix search
  } else if (query.length >= 1 && query.slice(-1).match(/\w/)) {
    query = query.slice(0, -1);
  }
  return query;
}