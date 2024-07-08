import { ThemedText } from "@/components/ThemedText";
import { useLocalSearchParams } from "expo-router";
import { View } from "react-native";

export default function Food() {
  const p = useLocalSearchParams();
  return (
    <View style={{ height: 300 }}>
      <ThemedText>Food: {JSON.stringify(p)}</ThemedText>
    </View>
  );
}
