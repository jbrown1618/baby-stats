import { ThemedText } from "@/components/ThemedText";
import { useLocalSearchParams } from "expo-router";
import { View } from "react-native";

export default function AddEvent() {
  const p = useLocalSearchParams();
  return (
    <View>
      <ThemedText>AddEvent: {JSON.stringify(p)}</ThemedText>
    </View>
  );
}
