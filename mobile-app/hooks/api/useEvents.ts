import { useQuery } from "@tanstack/react-query";
import { listEvents } from "./apiClient";
import { useGlobalSearchParams } from "expo-router";

export function useCurrentBabyEvents() {
  const p = useGlobalSearchParams<{ babyID: string }>();
  return useEvents(parseInt(p.babyID!));
}

export function useEvents(babyID: number) {
  return useQuery({
    queryKey: ["babies", babyID, "events"],
    queryFn: () => listEvents(babyID),
  });
}
