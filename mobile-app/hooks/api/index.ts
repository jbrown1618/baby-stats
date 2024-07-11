import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { listBabies, getBaby, listEvents, createEvent } from "./apiClient";
import { useGlobalSearchParams } from "expo-router";
import { NewEvent } from "@/types/api";

export function useBabies() {
  return useQuery({
    queryKey: ["babies"],
    queryFn: () => listBabies(),
  });
}

export function useBaby(babyID: number) {
  return useQuery({
    queryKey: ["babies", babyID],
    queryFn: () => getBaby(babyID),
  });
}

export function useCurrentBaby() {
  const p = useGlobalSearchParams<{ babyID: string }>();
  return useBaby(parseInt(p.babyID!));
}

export function useEvents(babyID: number) {
  return useQuery({
    queryKey: ["babies", babyID, "events"],
    queryFn: () => listEvents(babyID),
  });
}

export function useCreateEvent(babyID: number) {
  const client = useQueryClient();

  return useMutation({
    mutationFn: (newEvent: NewEvent) => createEvent(babyID, newEvent),
    onSuccess: () => {
      console.log("success");
      client.invalidateQueries({ queryKey: ["babies", babyID, "events"] });
    },
  });
}

export function useCurrentBabyEvents() {
  const p = useGlobalSearchParams<{ babyID: string }>();
  return useEvents(parseInt(p.babyID!));
}
