import { ThemedText } from "@/components/ThemedText";
import { useLocalSearchParams } from "expo-router";
import { View } from "react-native";

export default function Sleep() {
  const p = useLocalSearchParams();
  return (
    <View>
      <ThemedText>Sleep: {JSON.stringify(p)}</ThemedText>
    </View>
  );
}
