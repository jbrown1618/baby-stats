import { ThemedText } from "@/components/ThemedText";
import { useCurrentBabyEvents } from "@/hooks/api";
import { View } from "react-native";

export default function Timeline() {
  const { data: events } = useCurrentBabyEvents();
  return (
    <View>
      <ThemedText>Timeline: {JSON.stringify(events, null, 2)}</ThemedText>
    </View>
  );
}
