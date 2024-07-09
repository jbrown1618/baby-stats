import { useQuery } from "@tanstack/react-query";
import { getBaby } from "./apiClient";
import { useGlobalSearchParams } from "expo-router";

export function useCurrentBaby() {
  const p = useGlobalSearchParams<{ babyID: string }>();
  return useBaby(parseInt(p.babyID!));
}

export function useBaby(babyID: number) {
  return useQuery({
    queryKey: ["babies", babyID],
    queryFn: () => getBaby(babyID),
  });
}
