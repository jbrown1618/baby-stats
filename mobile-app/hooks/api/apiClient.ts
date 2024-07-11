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

async function makeRequest<T, B = void>(url: string, body?: B): Promise<T> {
  const res = await fetch(process.env.EXPO_PUBLIC_SERVER_URL + url, {
    method: body ? "POST" : "GET",
    body: body ? JSON.stringify(body) : undefined,
  });

  const json = await res.json();
  return json as T;
}
