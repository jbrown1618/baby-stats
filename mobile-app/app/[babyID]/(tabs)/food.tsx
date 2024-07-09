import { ThemedText } from "@/components/ThemedText";
import { useCurrentBaby } from "@/hooks/api";
import { View } from "react-native";

export default function Food() {
  const { data: baby } = useCurrentBaby();
  return (
    <View>
      <ThemedText>Food: {JSON.stringify(baby)}</ThemedText>
    </View>
  );
}
