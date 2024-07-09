export interface Baby {
  id: number;
  userId: number;
  name: string;
  birthDate: string;
}

export async function listBabies(): Promise<Baby[]> {
  return makeRequest<Baby[]>(`/babies`);
}

export async function getBaby(babyID: number): Promise<Baby> {
  return makeRequest<Baby>(`/babies/${babyID}`);
}

async function makeRequest<T>(url: string): Promise<T> {
  const res = await fetch(process.env.EXPO_PUBLIC_SERVER_URL + url);
  const json = await res.json();
  return json as T;
}
