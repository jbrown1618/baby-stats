import { ThemedText } from "@/components/ThemedText";
import { useCurrentBaby } from "@/hooks/api";
import { router } from "expo-router";
import { View } from "react-native";

export default function AddEvent() {
  const { data: baby } = useCurrentBaby();
  return (
    <View>
      <ThemedText>AddEvent: {JSON.stringify(baby)}</ThemedText>
      <ThemedText onPress={() => router.back()}>Dismiss</ThemedText>
    </View>
  );
}
