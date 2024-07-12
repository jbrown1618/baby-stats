import { Baby, NewEvent, Event } from "@/types/api";

export async function listBabies(): Promise<Baby[]> {
  return makeRequest<Baby[]>(`/babies`);
}

export async function getBaby(babyID: number): Promise<Baby> {
  return makeRequest<Baby>(`/babies/${babyID}`);
}

export async function listEvents(babyID: number): Promise<Event[]> {
  return makeRequest<Event[]>(`/babies/${babyID}/events`);
}

export async function createEvent(
  babyID: number,
  event: NewEvent
): Promise<Event> {
  return makeRequest<Event, NewEvent>(`/babies/${babyID}/events`, event);
}

async function makeRequest<T, B = void>(url: string, payload?: B): Promise<T> {
  const method = payload ? "POST" : "GET";
  const body = payload ? JSON.stringify(payload) : undefined;

  console.group(`${method} ${url}`);
  if (body) {
    console.debug("payload: ", body);
  }

  const res = await fetch(process.env.EXPO_PUBLIC_SERVER_URL + url, {
    method,
    body,
  });
  console.debug("status: ", res.status, res.statusText);

  const json = await res.json();
  console.debug("response: ", json);

  console.groupEnd();
  return json as T;
}
