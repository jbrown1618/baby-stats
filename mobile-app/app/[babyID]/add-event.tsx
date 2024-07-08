import { ThemedText } from "@/components/ThemedText";
import { useLocalSearchParams, router } from "expo-router";
import { View } from "react-native";

export default function AddEvent() {
  const p = useLocalSearchParams();
  return (
    <View>
      <ThemedText>AddEvent: {JSON.stringify(p)}</ThemedText>
      <ThemedText onPress={() => router.back()}>Dismiss</ThemedText>
    </View>
  );
}
