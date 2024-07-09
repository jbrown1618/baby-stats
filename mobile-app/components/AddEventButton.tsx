import { Link } from "expo-router";
import { ThemedText } from "./ThemedText";
import { useCurrentBaby } from "@/hooks/api";

export function AddEventButton() {
  const { data: baby, isLoading, isError } = useCurrentBaby();

  if (isLoading || isError || !baby) return null;

  return (
    <Link
      href={`/${baby!.id}/add-event`}
      style={{
        position: "absolute",
        right: 10,
        bottom: 60,
        backgroundColor: "red",
        padding: 10,
      }}
    >
      <ThemedText style={{ color: "white" }}>Add Event</ThemedText>
    </Link>
  );
}
