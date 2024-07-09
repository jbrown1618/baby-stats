import { useQuery } from "@tanstack/react-query";
import { listBabies } from "./apiClient";

export function useBabies() {
  return useQuery({
    queryKey: ["babies"],
    queryFn: () => listBabies(),
  });
}
